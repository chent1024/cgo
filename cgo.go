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
func New(conf *CgoConfig) (Cgo *gin.Engine) {
	// load config
	LoadConfig(conf.ConfigPath)

	// set timezone
	time.LoadLocation(Config.App.Timezone)

	// init gin
	gin.SetMode(Config.App.Mode)

	// log to file
	LogToFile()

	// new gin
	Cgo = gin.Default()
	funcs := template.FuncMap{
		"Unescaped": Unescaped,
	}
	for k, v := range conf.TplFuncMap {
		funcs[k] = v
	}
	Cgo.SetFuncMap(funcs)
	Cgo.LoadHTMLGlob(Config.Tpl.Path)
	Cgo.Routes()

	// init mysql
	NewMysql()

	return
}
