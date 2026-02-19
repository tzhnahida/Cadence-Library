package main

import (
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/alexbrainman/odbc"
)

func saveToAccess(data *AiResponse) (int, error) {
	connStr := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};DBQ=%s;", GlobalConfig.DbFile)
	db, err := sql.Open("odbc", connStr)
	if err != nil { return 0, err }
	defer db.Close()

	// 使用 Val 修正字符串排序问题
	var maxID sql.NullInt64
	queryMax := fmt.Sprintf("SELECT MAX(Val(Part_ID)) FROM [%s]", data.TableName)
	db.QueryRow(queryMax).Scan(&maxID)
	
	nextID := 1
	if maxID.Valid { nextID = int(maxID.Int64) + 1 }

	var cols, placeholders []string
	var vals []interface{}
	
	cols = append(cols, "[Part_ID]")
	placeholders = append(placeholders, "?")
	vals = append(vals, fmt.Sprintf("%d", nextID))

	for k, v := range data.Fields {
		if strings.EqualFold(k, "Part_ID") { continue }
		cols = append(cols, "["+k+"]")
		placeholders = append(placeholders, "?")
		vals = append(vals, v)
	}

	query := fmt.Sprintf("INSERT INTO [%s] (%s) VALUES (%s)", 
		data.TableName, strings.Join(cols, ", "), strings.Join(placeholders, ", "))
	
	_, err = db.Exec(query, vals...)
	return nextID, err
}