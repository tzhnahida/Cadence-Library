package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type AiResponse struct {
	TableName string                 `json:"table_name"`
	Fields    map[string]interface{} `json:"fields"`
}

type AiClient struct {
	Client  *resty.Client
	History []map[string]string
}

func NewAiClient() *AiClient {
	return &AiClient{
		Client: resty.New().SetAuthToken(GlobalConfig.ApiKey).SetBaseURL(GlobalConfig.BaseUrl),
		History: []map[string]string{
			{"role": "system", "content": GlobalConfig.SystemTag},
		},
	}
}

func (a *AiClient) Ask(content string) (*AiResponse, error) {
	a.History = append(a.History, map[string]string{"role": "user", "content": content})
	
	resp, err := a.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": GlobalConfig.AiModel,
			"messages": a.History,
			"response_format": map[string]string{"type": "json_object"},
		}).Post("/chat/completions")

	if err != nil { return nil, err }
	
	var res struct { Choices []struct { Message struct { Content string } } }
	json.Unmarshal(resp.Body(), &res)
	
	a.History = append(a.History, map[string]string{"role": "assistant", "content": res.Choices[0].Message.Content})

	var final AiResponse
	err = json.Unmarshal([]byte(res.Choices[0].Message.Content), &final)
	return &final, err
}