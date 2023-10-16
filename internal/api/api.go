package api

import (
	_ "account-transaction-api/docs"
	"account-transaction-api/internal/api/controllers"
	"account-transaction-api/internal/config"
	"account-transaction-api/internal/db"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {

	echo := echo.New()
	db := db.Init(cfg)

	echo.Use(echozap.ZapLogger(zap.L()))

	echo.Use(middleware.Recover())
	echo.Use(middleware.CORS())

	echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// health
	healthController := controllers.NewHealthController(echo, db)
	echo.GET("/", healthController.HealthCheck)
	echo.GET("/health", healthController.HealthCheck)

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
