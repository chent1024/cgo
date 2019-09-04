package cgo

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Debug
func Debug(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	t := time.Now().Format("2006/01/02 - 15:04:05")
	format = fmt.Sprintf("[Cgo] %v | "+format, t)
	fmt.Fprintf(gin.DefaultWriter, format, values...)
	fmt.Printf(format, values...)
}

// Write gin log to file
func NewLog() {
	cfg := Config.Log
	if !cfg.SaveLogs {
		return
	}
	os.MkdirAll(cfg.Path, os.ModePerm)

	gin.DisableConsoleColor()

	now := time.Now()

	// remove history log files
	daysAgo := now.AddDate(0, 0, -cfg.LogDays).Format("20060102")
	oldLogFile := fmt.Sprintf(cfg.Path+"/cgo_%s.log", daysAgo)
	os.Remove(oldLogFile)

	// write logs by day
	logFile := fmt.Sprintf(cfg.Path+"/cgo_%s.log", now.Format("20060102"))
	f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

}
