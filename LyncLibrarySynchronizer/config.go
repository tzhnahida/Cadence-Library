package main

import (
	"flag"
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

func LoadConfig() {
	// 1. 定义命令行参数 -c，默认值为当前目录下的 config.toml
	configPath := flag.String("c", "config.toml", "指定 config.toml 的完整路径")
	flag.Parse()

	log.Printf("📂 正在加载配置文件: %s", *configPath)

	// 2. 解码指定的 TOML 文件
	if _, err := toml.DecodeFile(*configPath, &GlobalConfig); err != nil {
		log.Fatalf("❌ 无法加载配置文件 [%s]: %v\n提示: 请检查路径是否正确或文件格式是否为 UTF-8", *configPath, err)
	}
}