package main

import (
	"embed"
	"log"
	"os"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed frontend/*
var assets embed.FS

func main() {
	configPath := "config.toml"
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config.toml 未找到，请将 config.toml 模板放在程序目录下并填写配置")
	}
	LoadConfig(configPath)

	app := application.New(application.Options{
		Name: "PCB Library 助手",
		Assets: application.AssetOptions{
			Handler: application.BundledAssetFileServer(assets),
		},
		Services: []application.Service{
			application.NewService(NewAppService()),
			application.NewService(NewConfigService()),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		URL:             "/",
		Title:           "PCB Library 助手",
		Width:           1100,
		Height:          750,
		DevToolsEnabled: true,
	})

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
