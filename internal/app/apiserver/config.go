package apiserver

import "http-rest-api/internal/app/store"

// Config ...
type Config struct {
	BindAddress string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}