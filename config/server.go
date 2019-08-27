package config

// server服务
type ServerConfig struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}
