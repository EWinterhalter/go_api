package apiserver

import "github.com/EWinterhalter/go_api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_add"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
