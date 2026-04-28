package main

import (
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/alexbrainman/odbc"
)

type TableSchema struct {
	Name    string   `json:"name"`
	Columns []string `json:"columns"`
}

type TableStat struct {
	Name    string `json:"name"`
	Records int    `json:"records"`
}

func openDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};DBQ=%s;", GlobalConfig.DbFile)
	return sql.Open("odbc", connStr)
}

func GetTableSchemas() ([]TableSchema, error) {
	tables := []string{
		"IC_Params", "Resistor_Params", "Crystal_Params", "LED_Params",
		"Diode_Params", "Inductor_Params", "FerriteBead_Params",
		"Connector_Params", "Miscellaneous_Params",
	}

	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var schemas []TableSchema
	for _, t := range tables {
		rows, err := db.Query(fmt.Sprintf("SELECT TOP 1 * FROM [%s]", t))
		if err != nil {
			continue
		}
		cols, err := rows.Columns()
		rows.Close()
		if err != nil {
			continue
		}
		schemas = append(schemas, TableSchema{Name: t, Columns: cols})
	}
	return schemas, nil
}

func GetDBStats() ([]TableStat, error) {
	tables := []string{
		"IC_Params", "Resistor_Params", "Crystal_Params", "LED_Params",
		"Diode_Params", "Inductor_Params", "FerriteBead_Params",
		"Connector_Params", "Miscellaneous_Params",
	}

	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var stats []TableStat
	for _, t := range tables {
		var count int
		err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM [%s]", t)).Scan(&count)
		if err != nil {
			continue
		}
		stats = append(stats, TableStat{Name: t, Records: count})
	}
	return stats, nil
}

func BuildDBContext() string {
	db, err := openDB()
	if err != nil {
		return ""
	}
	defer db.Close()

	tables := []string{
		"IC_Params", "Resistor_Params", "Crystal_Params", "LED_Params",
		"Diode_Params", "Inductor_Params", "FerriteBead_Params",
		"Connector_Params", "Miscellaneous_Params",
	}

	var sb strings.Builder
	sb.WriteString("=== 现有数据库记录参考，请严格沿用以下命名惯例 ===\n")

	for _, table := range tables {
		var count int
		db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM [%s]", table)).Scan(&count)
		if count == 0 {
			continue
		}

		sb.WriteString(fmt.Sprintf("\n【%s】(共%d条)", table, count))

		rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT TOP 8 Category FROM [%s] WHERE Category IS NOT NULL AND Category <> '' AND Category <> '-'", table))
		if err == nil {
			var cats []string
			for rows.Next() {
				var c string
				rows.Scan(&c)
				cats = append(cats, c)
			}
			rows.Close()
			if len(cats) > 0 {
				sb.WriteString(fmt.Sprintf("\n  已有分类: %s", strings.Join(cats, ", ")))
			}
		}

		rows, err = db.Query(fmt.Sprintf("SELECT DISTINCT TOP 8 Symbol_Name FROM [%s] WHERE Symbol_Name IS NOT NULL AND Symbol_Name <> '' AND Symbol_Name <> '-'", table))
		if err == nil {
			var syms []string
			for rows.Next() {
				var s string
				rows.Scan(&s)
				syms = append(syms, s)
			}
			rows.Close()
			if len(syms) > 0 {
				sb.WriteString(fmt.Sprintf("\n  已有Symbol: %s", strings.Join(syms, ", ")))
			}
		}

		rows, err = db.Query(fmt.Sprintf("SELECT DISTINCT TOP 8 Footprint_Name FROM [%s] WHERE Footprint_Name IS NOT NULL AND Footprint_Name <> '' AND Footprint_Name <> '-'", table))
		if err == nil {
			var fps []string
			for rows.Next() {
				var f string
				rows.Scan(&f)
				fps = append(fps, f)
			}
			rows.Close()
			if len(fps) > 0 {
				sb.WriteString(fmt.Sprintf("\n  已有封装: %s", strings.Join(fps, ", ")))
			}
		}
	}

	return sb.String()
}

func saveToAccess(data *AiResponse) (int, error) {
	db, err := openDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var maxID sql.NullInt64
	queryMax := fmt.Sprintf("SELECT MAX(Val(Part_ID)) FROM [%s]", data.TableName)
	db.QueryRow(queryMax).Scan(&maxID)

	nextID := 1
	if maxID.Valid {
		nextID = int(maxID.Int64) + 1
	}

	var cols, placeholders []string
	var vals []interface{}

	cols = append(cols, "[Part_ID]")
	placeholders = append(placeholders, "?")
	vals = append(vals, fmt.Sprintf("%d", nextID))

	for k, v := range data.Fields {
		if strings.EqualFold(k, "Part_ID") {
			continue
		}
		cols = append(cols, "["+k+"]")
		placeholders = append(placeholders, "?")
		vals = append(vals, v)
	}

	query := fmt.Sprintf("INSERT INTO [%s] (%s) VALUES (%s)",
		data.TableName, strings.Join(cols, ", "), strings.Join(placeholders, ", "))

	_, err = db.Exec(query, vals...)
	return nextID, err
}
