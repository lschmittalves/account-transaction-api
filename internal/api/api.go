package api

import (
	_ "account-transaction-api/docs"
	"account-transaction-api/internal/api/controllers"
	"account-transaction-api/internal/cache"
	"account-transaction-api/internal/config"
	"account-transaction-api/internal/db"
	"account-transaction-api/internal/repositories"
	accountService "account-transaction-api/internal/services/account"
	transactionService "account-transaction-api/internal/services/transactions"
	"github.com/brpaz/echozap"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

func Run(cfg *config.Config) {

	e := echo.New()
	db := db.Init(cfg)
	c := cache.Init(cfg)

	e.Use(echozap.ZapLogger(zap.L()))

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// metrics
	e.Use(echoprometheus.NewMiddleware("api"))
	e.GET("/metrics", echoprometheus.NewHandler())

	// health
	healthController := controllers.NewHealthController(e, db)
	e.GET("/", healthController.HealthCheck)
	e.GET("/health", healthController.HealthCheck)

	// accounts
	accountRepository := repositories.NewAccountRepository(db)
	accountsController := controllers.NewAccountsController(e, accountService.NewAccountService(accountRepository, accountRepository), accountRepository)
	e.POST("/accounts", accountsController.Post)
	e.GET("/accounts/:id", accountsController.Get)

	// transactions
	operationsRepository := repositories.NewOperationTypeRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db, c)
	transactionController := controllers.NewTransactionsController(e, transactionService.NewTransactionService(accountRepository, operationsRepository, transactionRepository))
	e.POST("/transactions", transactionController.Post)

	err := e.Start(":" + cfg.HTTP.Port)
	if err != nil {
		panic("Port already used")
	}
}
