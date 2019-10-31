package cgo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
	reset     = string([]byte{27, 91, 48, 109})
)

var Logger *logrus.Logger

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

	format = fmt.Sprintf("[Cgo] %v | "+format, timeFormat(time.Now()))
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

	var logFile string
	if Config.Log.Single {
		logFile = fmt.Sprintf(cfg.Path+"/%s.log", cfg.LogName)
	} else {
		// write logs by day
		logFile = fmt.Sprintf(cfg.Path+"/%s_%s.log", cfg.LogName, time.Now().Format("20060102"))
	}
	f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
}

func LogrusMiddleware() gin.HandlerFunc {
	logger := NewLogrus()
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.SetLevel(logrus.ErrorLevel)
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if logger != nil {
					stack := stack(3)
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					headers := strings.Split(string(httpRequest), "\r\n")
					for idx, header := range headers {
						current := strings.Split(header, ":")
						if current[0] == "Authorization" {
							headers[idx] = current[0] + ": *"
						}
					}
					if brokenPipe {
						logger.Errorf("%s\n%s%s", err, string(httpRequest), reset)
					} else if gin.IsDebugging() {
						logger.Panicf("[Recovery] %s panic recovered:\n%s\n%s\n%s%s",
							timeFormat(time.Now()), strings.Join(headers, "\r\n"), err, stack, reset)
					} else {
						logger.Panicf("[Recovery] %s panic recovered:\n%s\n%s%s",
							timeFormat(time.Now()), err, stack, reset)
					}
				}

				// If the connection is dead, we can't write a status to it.
				if brokenPipe {
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()

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
		logger.Infof("| %3d | %13v | %15s | %s | %s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func NewLogrus() *logrus.Logger {
	logger := logrus.New()

	os.MkdirAll(Config.Log.Path, os.ModePerm)
	var logFile string
	if Config.Log.Single {
		logFile = fmt.Sprintf(Config.Log.Path+"/%s.log", Config.Log.LogName)
	} else {
		logFile = fmt.Sprintf(Config.Log.Path+"/%s_%s.log", Config.Log.LogName, time.Now().Format("20060102"))
	}
	src, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		Loginfo("log middleware err %v", err)
	}

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.Out = src

	// 供外部调用
	Logger = logger

	logger.SetLevel(Config.Log.Level)
	return logger
}

// stack returns a nicely formatted stack frame, skipping skip frames.
func stack(skip int) []byte {
	buf := new(bytes.Buffer) // the returned data
	// As we loop, we open files and read them. These variables record the currently
	// loaded file.
	var lines [][]byte
	var lastFile string
	for i := skip; ; i++ { // Skip the expected number of frames
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		// Print this much at least.  If we can't find the source, it won't show.
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
		if file != lastFile {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			lines = bytes.Split(data, []byte{'\n'})
			lastFile = file
		}
		fmt.Fprintf(buf, "\t%s: %s\n", function(pc), source(lines, line))
	}
	return buf.Bytes()
}

// source returns a space-trimmed slice of the n'th line.
func source(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return dunno
	}
	return bytes.TrimSpace(lines[n])
}

// function returns, if possible, the name of the function containing the PC.
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//	runtime/debug.*T·ptrmethod
	// and want
	//	*T.ptrmethod
	// Also the package path might contains dot (e.g. code.google.com/...),
	// so first eliminate the path prefix
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}

func timeFormat(t time.Time) string {
	var timeString = t.Format("2006/01/02 - 15:04:05")
	return timeString
}
