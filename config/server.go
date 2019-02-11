package config

// server服务
type ServerConfig struct {
	Address      string `ini:"address"`
	ReadTimeout  int    `ini:"read_timeout"`
	WriteTimeout int    `ini:"write_timeout"`
	IdleTimeout  int    `ini:"idle_timeout"`
}
