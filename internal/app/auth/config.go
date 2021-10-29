package auth

type Config struct {
	TokenHeader string `toml:"token_header"`
	AuthSecret string `toml:"auth_secret"`
}

func NewConfig() *Config {
	return &Config{
		TokenHeader: "Token",
		AuthSecret: "secret",
	}
}
