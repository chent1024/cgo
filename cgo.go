package cgo

import (
	"html/template"
	"net/http"
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
	// NewLogWriter()
	g = gin.Default()
	g.Use(LogrusMiddleware())
	NewTemplate(g, conf)
	NewMysql()

	return g
}

func Run(gin *gin.Engine) {
	server := &http.Server{
		Addr:         Config.Server.Address,
		Handler:      gin,
		ReadTimeout:  Config.Server.ReadTimeout * time.Second,
		WriteTimeout: Config.Server.WriteTimeout * time.Second,
		IdleTimeout:  Config.Server.IdleTimeout * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		Logger.Panic("Listen err ", err)
	}
}
