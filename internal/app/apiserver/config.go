package apiserver

// Config ...
type Config struct {
	BindAddress string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel: "debug",
	}
}