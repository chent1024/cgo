package cgo

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type CgoConfig struct {
	ConfigPath string
	TplFuncMap template.FuncMap
}

// New cgo with gin
func New(conf *CgoConfig) (g *gin.Engine) {
	// load config
	NewConfig(conf.ConfigPath)
	time.LoadLocation(Config.App.Timezone)

	// new gin
	gin.SetMode(Config.App.Mode)
	g = gin.Default()
	g.Use(LogrusMiddleware())

	// is use template
	if Config.Tpl.Enable && Config.Tpl.Path != "" {
		NewTemplate(g, conf)
	}
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

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Listen err: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Err: ", err)
	}

	log.Println("Server Shutdown Success")

}
