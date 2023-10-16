package helpers

import (
	"account-transaction-api/internal/api"
	"account-transaction-api/internal/config"
	"github.com/labstack/echo/v4"
)

func NewServer() *api.Server {
	s := &api.Server{
		Echo:   echo.New(),
		DB:     Init(),
		Config: config.NewConfig(),
	}

	return s
}
