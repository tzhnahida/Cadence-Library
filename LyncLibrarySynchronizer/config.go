package main

import (
	"log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	ApiKey    string `toml:"api_key"`
	BaseUrl   string `toml:"base_url"`
	AiModel   string `toml:"ai_model"`
	DbFile    string `toml:"db_file"`
	SystemTag string `toml:"system_tag"`
}

var GlobalConfig Config

func LoadConfig(path string) {
	log.Printf("📂 正在加载配置文件: %s", path)
	if _, err := toml.DecodeFile(path, &GlobalConfig); err != nil {
		log.Fatalf("❌ 无法加载配置文件 [%s]: %v\n提示: 请检查路径是否正确或文件格式是否为 UTF-8", path, err)
	}
}
