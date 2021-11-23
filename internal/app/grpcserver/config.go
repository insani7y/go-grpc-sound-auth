package grpcserver

type Config struct {
	BindGRPCAddr string `toml:"bind_grpc_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindGRPCAddr: ":9000",
	}
}
