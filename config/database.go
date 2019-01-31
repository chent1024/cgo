package config

type DbConfig struct {
	Debug           bool   `ini:"debug"`
	Host            string `ini:"host"`
	Charset         string `ini:"charset"`
	Username        string `ini:"username"`
	Password        string `ini:"password"`
	Database        string `ini:"database"`
	TablePrefix     string `ini:"table_prefix"`
	MaxIdleConns    int    `ini:"max_idle_conns"`
	MaxOpenConns    int    `ini:"max_open_cons"`
	ConnMaxLifeTime int    `ini:"conn_max_life_time"`
}
