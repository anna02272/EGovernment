package config

import (
	"os"
)

type Config struct {
	SecretKey string
}

func LoadConfig() *Config {
	return &Config{
		SecretKey: os.Getenv("SECRET_KEY"),
	}

}
