package config

import (
	"os"
)

var defaultENV = map[string]string{
	"PORT":      ":8000",
	"DB":        "postgresql://xxx:xxx@localhost/todo",
	"LOG_LEVEL": "debug",
}

type ENVConfig struct {
	Port     string
	DB       string
	LogLevel string
}

func NewENVConfig() *ENVConfig {
	env := map[string]string{}
	for k, v := range defaultENV {
		env[k] = os.Getenv(k)
		if env[k] == "" {
			env[k] = v
		}
	}
	return &ENVConfig{
		Port:     env["PORT"],
		DB:       env["DB"],
		LogLevel: env["LOG_LEVEL"],
	}
}
