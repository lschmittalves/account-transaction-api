package helpers

import (
	"account-transaction-api/config"
	"account-transaction-api/server"
	"github.com/labstack/echo/v4"
)

func NewServer() *server.Server {
	s := &server.Server{
		Echo:   echo.New(),
		DB:     Init(),
		Config: config.NewConfig(),
	}

	return s
}
