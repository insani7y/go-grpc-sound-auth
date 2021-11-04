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
		MFCCUrl: "http://127.0.0.1:8000/mfcc/",
	}
}
