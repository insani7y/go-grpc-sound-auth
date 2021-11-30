package auth

type Config struct {
	TokenHeader string `toml:"token_header"`
	AuthSecret string `toml:"auth_secret"`
	MFCCUrl string `toml:"mfcc_url"`
}

func NewConfig() *Config {
	return &Config{
		TokenHeader: "Token",
		AuthSecret: "secret",
		MFCCUrl: "http://python-backend:8000/py/mfcc/",
	}
}
