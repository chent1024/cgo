package config

// authorization 认证配置
type AuthorizationConfig struct {
	EncryptKey   string
	EncryptIv    string
	ExpectRouter []string
}
