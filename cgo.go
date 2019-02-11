package cgo

import (
	"github.com/gin-gonic/gin"
)

// New cgo with gin
func New(configPath string) (Cgo *gin.Engine) {
	// load config
	LoadConfig(configPath)

	// init gin
	gin.SetMode(Config.App.Mode)

	// log to file
	LogToFile()

	Cgo = gin.Default()
	// load tpl
	Cgo.LoadHTMLGlob(Config.Tpl.Path)

	// init mysql
	InitMysql()

	return
}
