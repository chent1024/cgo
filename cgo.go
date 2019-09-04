package cgo

import (
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

type CgoConfig struct {
	ConfigPath string
	TplFuncMap template.FuncMap
}

// New cgo with gin
func New(conf *CgoConfig) (g *gin.Engine) {
	NewConfig(conf.ConfigPath)
	time.LoadLocation(Config.App.Timezone)

	gin.SetMode(Config.App.Mode)
	// log to file
	NewLog()
	g = gin.Default()
	g.Routes()
	NewTemplate(g, conf)
	NewMysql()

	return
}
