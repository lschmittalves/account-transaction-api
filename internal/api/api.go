package api

import (
	"account-transaction-api/internal/api/controllers"
	"account-transaction-api/internal/config"
	"account-transaction-api/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Run(cfg *config.Config) {

	echo := echo.New()
	db := db.Init(cfg)
	echo.Use(middleware.Logger())

	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// accounts
	accountsController := controllers.NewAccountsController(echo, db)
	echo.POST("/accounts", accountsController.Post)
	echo.GET("/accounts/:id", accountsController.Get)

	// transactions
	//transactionController := controllers.NewTransactionController(server)
	//server.Echo.POST("/transactions", transactionController.Post)
	//server.Echo.GET("/transactions/:id", transactionController.Get)

	err := echo.Start(":" + cfg.HTTP.Port)
	if err != nil {
		panic("Port already used")
	}
}
