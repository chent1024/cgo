package config

// 数据库
type DbConfig struct {
	Debug           bool
	Host            string
	Charset         string
	Username        string
	Password        string
	Database        string
	TablePrefix     string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime int
}
