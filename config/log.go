package config

// 日志
type LogConfig struct {
	SaveLogs bool   `ini:"save_logs"`
	Path     string `ini:"path"`
	LogDays  int    `ini:"log_days"`
}
