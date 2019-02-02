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

	if err := cfg.Section("app").MapTo(&Config.App); err != nil {
		Debug("app config not found")
		return
	}

	if err := cfg.Section("log").MapTo(&Config.Log); err != nil {
		Debug("log config not found")
		return
	}

	if err := cfg.Section("server").MapTo(&Config.Server); err != nil {
		Debug("server config not found")
		return
	}

	if err := cfg.Section("database").MapTo(&Config.Db); err != nil {
		Debug("database config not found")
		return
	}

	if err := cfg.Section("template").MapTo(&Config.Tpl); err != nil {
		Debug("template config not found")
		return
	}

	return
}
