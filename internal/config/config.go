package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	DB    DBConfig
	CACHE CacheConfig
	HTTP  HTTPConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		zap.L().Error("Error loading .env file")
	}

	return &Config{
		CACHE: LoadCacheConfig(),
		DB:    LoadDBConfig(),
		HTTP:  LoadHTTPConfig(),
	}
}
