package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthController struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewHealthController(echo *echo.Echo, db *gorm.DB) *HealthController {
	return &HealthController{Echo: echo, DB: db}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (controller *HealthController) HealthCheck(c echo.Context) error {

	var databaseStatus = "Not-Connected"

	if err := controller.DB.DB().Ping(); err == nil {
		databaseStatus = "Connected"
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"server":   "Server is up and running",
		"database": databaseStatus,
	})
}
