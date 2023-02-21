package config

import (
	"context"
	"os"
)

type Config struct {
	Context  *context.Context
	BindAddr string
	LogLevel string
	DbUrl    string
}

func NewConfig(ctx *context.Context) *Config {
	return &Config{
		Context:  ctx,
		BindAddr: ":8888",
		LogLevel: "debug",
		DbUrl:    os.Getenv("DB_URL"),
	}
}