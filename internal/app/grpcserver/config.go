package grpcserver

type Config struct {
	BindGRPCAddr string `toml:"bind_grpc_addr"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindGRPCAddr: ":9000",
	}
}
