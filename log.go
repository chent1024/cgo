package cgo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
	"time"
)

func Debug(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Fprintf(os.Stderr, "[Cgo-debug] "+format, values...)
}

func LogToFile() {
	cfg := Config.Log
	if !cfg.SaveLogs {
		return
	}

	os.MkdirAll(cfg.Path, os.ModePerm)

	gin.DisableConsoleColor()

	now := time.Now()

	// remove history log files
	oldLogPath := fmt.Sprintf(cfg.Path+"/cgo_%s.log", now.AddDate(0, 0, -cfg.LogDays).Format("20060102"))
	os.Remove(oldLogPath)

	// write logs by day
	logPath := fmt.Sprintf(cfg.Path+"/cgo_%s.log", now.Format("20060102"))
	f, _ := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	gin.DefaultWriter = io.MultiWriter(f)

}
