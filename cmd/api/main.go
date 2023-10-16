package main

import (
	"account-transaction-api/docs"
	"account-transaction-api/internal/api"
	"account-transaction-api/internal/config"
	"fmt"
)

func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	api.Run(cfg)
}
