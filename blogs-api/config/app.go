package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

var Config = &App{}

type App struct {
	Page   Page   `ini:"page"`
	Mysql  MySQL  `ini:"mysql"`
	Server Server `ini:"server"`
	Debug  bool   `ini:"debug"`
	Redis  Redis  `ini:"redis"`
	Upload Upload `ini:"upload"`
}
type IniConfig struct {
	ConfigPath       string
	ConfigName       string
	RunModelName     string
	RunModel         string
	RunModelErrAllow bool // 允许不使用拓展配置文件
}

func InitConfig(config *IniConfig) {
	if config.ConfigName == "" {
		config.ConfigName = "app.ini"
	}
	// 读取配置文件
	load, err := ini.Load(filepath.Join(config.ConfigPath, config.ConfigName))
	if err != nil {
		log.Fatal("配置文件读取错误，err:", err)
		return
	}
	// 环境配置
	if config.RunModel == "" {
		config.RunModel = os.Getenv("RunModel")
		if config.RunModel == "" {
			config.RunModel = load.Section("").Key("defaultModel").String()
			if config.RunModel == "" {
				config.RunModel = "test"
			}
		}
	}
	if config.RunModelName == "" {
		config.RunModelName = fmt.Sprintf("app.%s.ini", config.RunModel)
	}
	if err := load.Append(filepath.Join(config.ConfigPath, config.RunModelName)); err != nil {
		if config.RunModelErrAllow {
			log.Println("环境配置文件读取错误，err:", err)
		} else {
			log.Fatal("环境配置文件读取错误，err:", err)
		}
	}
	// 映射配置
	if err := load.MapTo(Config); err != nil {
		log.Fatal("配置文件映射错误，err:", err)
		return
	}
}
