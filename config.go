package cgo

import (
	"github.com/chent1024/cgo/config"
	"github.com/go-ini/ini"
)

var Config struct {
	App    config.AppConfig
	Log    config.LogConfig
	Db     config.DbConfig
	Server config.ServerConfig
	Tpl    config.TemplateConfig
}

func LoadConfig(path string) {
	cfg, err := ini.Load(path)
	if err != nil {
		Debug(err.Error())
		return
	}

	cfg.Section("app").MapTo(&Config.App)
	cfg.Section("log").MapTo(&Config.Log)
	cfg.Section("server").MapTo(&Config.Server)
	cfg.Section("database").MapTo(&Config.Db)
	cfg.Section("template").MapTo(&Config.Tpl)

	return
}
