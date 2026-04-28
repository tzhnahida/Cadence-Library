package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigService struct{}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (c *ConfigService) GetConfig() Config {
	return GlobalConfig
}

func (c *ConfigService) UpdateConfig(cfg Config) error {
	cfg.SystemTag = stringsTrim(cfg.SystemTag)
	if cfg.ApiKey == "" {
		return fmt.Errorf("API Key 不能为空")
	}
	if cfg.DbFile == "" {
		return fmt.Errorf("数据库路径不能为空")
	}

	GlobalConfig = cfg

	f, err := os.Create("config.toml")
	if err != nil {
		return fmt.Errorf("无法写入配置文件: %w", err)
	}
	defer f.Close()

	if err := toml.NewEncoder(f).Encode(cfg); err != nil {
		return fmt.Errorf("编码配置失败: %w", err)
	}
	return nil
}

func (c *ConfigService) GetDBStats() ([]TableStat, error) {
	return GetDBStats()
}

func (c *ConfigService) PingDB() (bool, error) {
	db, err := openDB()
	if err != nil {
		return false, err
	}
	defer db.Close()
	return db.Ping() == nil, nil
}

func stringsTrim(s string) string {
	if len(s) > 0 && s[0] == '\n' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	return s
}
