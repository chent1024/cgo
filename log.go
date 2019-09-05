package cgo

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Debug
func Debug(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Fprintf(os.Stderr, "[Cgo-debug] "+format, values...)
}

// Log info to file
func Loginfo(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	format = fmt.Sprintf("[Cgo] %v | "+format, time.Now().Format("2006/01/02 - 15:04:05"))
	fmt.Fprintf(gin.DefaultWriter, format, values...)
}

// Custom log writer to file
func NewLogWriter() {
	cfg := Config.Log
	if !cfg.SaveLogs {
		return
	}
	os.MkdirAll(cfg.Path, os.ModePerm)

	gin.DisableConsoleColor()

	now := time.Now()
	// write logs by day
	logFile := fmt.Sprintf(cfg.Path+"/%s_%s.log", cfg.LogName, now.Format("20060102"))
	f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
}

func LogrusMiddleware() gin.HandlerFunc {
	logger := logrus.New()

	os.MkdirAll(Config.Log.Path, os.ModePerm)

	logFile := fmt.Sprintf(Config.Log.Path+"/%s_%s.log", Config.Log.LogName, time.Now().Format("20060102"))
	src, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		Loginfo("log middleware err %v", err)
	}

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.Out = src
	logger.SetLevel(Config.Log.Level)
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
