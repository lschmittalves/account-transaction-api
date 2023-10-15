package main

import (
	application "account-transaction-api"
	"account-transaction-api/config"
	"account-transaction-api/docs"
	"fmt"
)

// @BasePath /
func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	application.Start(cfg)
}
