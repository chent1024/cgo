package cgo

import (
	"github.com/chent1024/cgo/config"
	"github.com/gin-gonic/gin"
)

type Cgo struct {
	Gin *gin.Engine
}

func New(configPath string) (cgo *Cgo) {
	cgo = &Cgo{}

	// load config
	cfg := config.Config{
		Path: configPath,
	}
	cfg.New()

	// init gin
	gin.SetMode(config.Conf.App.Mode)
	cgo.Gin = gin.Default()

	// init mysql
	InitMysql()

	return
}
