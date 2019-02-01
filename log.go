package cgo

import (
	"fmt"
	"github.com/chent1024/cgo/config"
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
	cfg := config.Conf.Log
	if !cfg.SaveLogs {
		return
	}

	os.MkdirAll(cfg.Path, os.ModePerm)

	gin.DisableConsoleColor()

	// remove history log files
	oldLogPath := fmt.Sprintf(cfg.Path+"/cgo_%s.log", time.Now().AddDate(0, 0, -cfg.LogDays).Format("20060102"))
	os.Remove(oldLogPath)

	// write logs by day
	logPath := fmt.Sprintf(cfg.Path+"/cgo_%s.log", time.Now().Format("20060102"))
	f, _ := os.Create(logPath)
	gin.DefaultWriter = io.MultiWriter(f)

}
