package cgo

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/chent1024/cgo/config"
)

// define Config struct
var Config struct {
	App           config.AppConfig
	Log           config.LogConfig
	Db            config.DbConfig
	Server        config.ServerConfig
	Tpl           config.TemplateConfig
	Jwt           config.JwtConfig
	Redis         config.RedisConfig
	Authorization config.AuthorizationConfig
}

// new Config
func NewConfig(path string) {
	LoadConfig(path, &Config)
}

// Load config
func LoadConfig(path string, resp interface{}) {
	_, err := toml.DecodeFile(path, resp)
	if err != nil {
		panic(fmt.Sprintf("Load config failed, %#v", err.Error()))
	}
}
