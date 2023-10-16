package main

import (
	"account-transaction-api/internal/api"
	"account-transaction-api/internal/config"
	"go.uber.org/zap"
	"os"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
	os.Setenv("TZ", "Etc/UTC")
}

// @title Account Transactions API
// @version 1.0
// @description This is a sample of an account transactions repository.

// @contact.name Leonardo Alves
// @contact.url http://www.swagger.io/support
// @contact.email lschmittalves@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @schemes http

func main() {
	cfg := config.NewConfig()
	api.Run(cfg)
}
