package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	DB    DBConfig
	CACHE CacheConfig
	HTTP  HTTPConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		CACHE: LoadCacheConfig(),
		DB:    LoadDBConfig(),
		HTTP:  LoadHTTPConfig(),
	}
}
