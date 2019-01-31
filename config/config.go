package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

type Config struct {
	Path   string
	App    AppConfig
	Db     DbConfig
	Server ServerConfig
	Tpl    TemplateConfig
}

type AppConfig struct {
	Mode string `ini:"mode"`
}

var Conf Config

func (c *Config) New() {
	cfg, err := ini.Load(c.Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.Section("app").MapTo(&Conf.App)
	cfg.Section("server").MapTo(&Conf.Server)
	cfg.Section("database").MapTo(&Conf.Db)
	cfg.Section("template").MapTo(&Conf.Tpl)

	return
}
