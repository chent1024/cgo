package cgo

import (
	"github.com/BurntSushi/toml"
	"github.com/chent1024/cgo/config"
)

// define Config struct
var Config struct {
	App    config.AppConfig
	Log    config.LogConfig
	Db     config.DbConfig
	Server config.ServerConfig
	Tpl    config.TemplateConfig
	Jwt    config.JwtConfig
	Redis  config.RedisConfig
}

// load config to &Config
func NewConfig(path string) {
	_, err := toml.DecodeFile(path, &Config)
	if err != nil {
		Loginfo("load config faild, %s", err.Error())
		return
	}

	return
}
