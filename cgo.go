package cgo

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

type CgoConfig struct {
	ConfigPath string
	TplFuncMap template.FuncMap
}

// New cgo with gin
func New(conf *CgoConfig) (g *gin.Engine) {
	// load config
	LoadConfig(conf.ConfigPath)

	// set timezone
	time.LoadLocation(Config.App.Timezone)

	// init gin
	gin.SetMode(Config.App.Mode)

	// log to file
	LogToFile()

	// new gin
	g = gin.Default()
	funcs := template.FuncMap{
		"Unescaped": Unescaped,
	}
	for k, v := range conf.TplFuncMap {
		funcs[k] = v
	}
	g.SetFuncMap(funcs)
	g.LoadHTMLGlob(Config.Tpl.Path)
	g.Routes()

	// init mysql
	NewMysql()

	return
}
