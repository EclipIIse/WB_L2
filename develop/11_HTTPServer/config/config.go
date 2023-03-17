package config

import "github.com/EclipIIse/WB_L2/develop/11_HTTPServer/store"

type Config struct {
	Port  string
	Store *store.Config
}

func NewDefConfig() *Config {
	return &Config{
		Port:  ":8080",
		Store: store.NewDbConfig(),
	}
}
