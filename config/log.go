package config

import "github.com/sirupsen/logrus"

// 日志
type LogConfig struct {
	SaveLogs bool
	Single   bool
	Level    logrus.Level
	Path     string
	LogName  string
}
