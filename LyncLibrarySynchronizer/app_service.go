package main

import (
	"fmt"
	"strings"
)

type AppService struct {
	ai *AiClient
}

func NewAppService() *AppService {
	return &AppService{
		ai: NewAiClient(),
	}
}

type AnalysisResult struct {
	LCSCID     string                 `json:"lcsc_id"`
	TableName  string                 `json:"table_name"`
	Fields     map[string]interface{} `json:"fields"`
	TableNames []string               `json:"table_names"`
	RawData    string                 `json:"raw_data,omitempty"`
}

type SaveResult struct {
	PartID    int    `json:"part_id"`
	SystemTag string `json:"system_tag"`
	TableName string `json:"table_name"`
}

func (a *AppService) AnalyzeLCSC(url string) (*AnalysisResult, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return nil, fmt.Errorf("URL 不能为空")
	}

	raw, id, err := fetchLCSC(url)
	if err != nil {
		return nil, fmt.Errorf("抓取失败: %w", err)
	}

	schemas, _ := GetTableSchemas()

	var schemaInfo strings.Builder
	schemaInfo.WriteString("=== 数据库表结构 ===\n")
	tableNames := make([]string, 0)
	for _, s := range schemas {
		tableNames = append(tableNames, s.Name)
		schemaInfo.WriteString(fmt.Sprintf("[%s] 字段: %s\n", s.Name, strings.Join(s.Columns, ", ")))
	}

	dbCtx := BuildDBContext()

	prompt := fmt.Sprintf("编号: %s\n\n%s\n\n%s\n\n内容:\n%s", id, schemaInfo.String(), dbCtx, raw)
	parsed, err := a.ai.Ask(prompt)
	if err != nil {
		return nil, fmt.Errorf("AI 分析失败: %w", err)
	}

	return &AnalysisResult{
		LCSCID:     id,
		TableName:  parsed.TableName,
		Fields:     parsed.Fields,
		TableNames: tableNames,
		RawData:    raw,
	}, nil
}

func (a *AppService) SaveToDatabase(data *AnalysisResult) (*SaveResult, error) {
	if data == nil || len(data.Fields) == 0 {
		return nil, fmt.Errorf("无数据可保存")
	}

	parsed := &AiResponse{
		TableName: data.TableName,
		Fields:    data.Fields,
	}

	newID, err := saveToAccess(parsed)
	if err != nil {
		return nil, fmt.Errorf("写入失败: %w", err)
	}

	return &SaveResult{
		PartID:    newID,
		SystemTag: GlobalConfig.SystemTag,
		TableName: data.TableName,
	}, nil
}

func (a *AppService) ClearHistory() {
	a.ai = NewAiClient()
}
