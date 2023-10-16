package controllers

import (
	"account-transaction-api/internal/api/requests"
	"account-transaction-api/internal/api/responses"
	transactionService "account-transaction-api/internal/services/transactions"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionsController struct {
	echo    *echo.Echo
	service transactionService.Service
}

func NewTransactionsController(echo *echo.Echo, service transactionService.Service) *TransactionsController {
	return &TransactionsController{echo: echo, service: service}
}

// Post godoc
// @Summary Registers a new transaction
// @Description Registers a new transaction passing the account id, operation id and amount, the amount can change to a negative value in the database if the business logic decides that it`s a debit operation. The amount value must be sent in base 100 in order to avoid float point precision issues
// @ID transaction-create
// @Tags Transaction Actions
// @Accept json
// @Produce json
// @Param params body requests.RegisterTransactionRequest true "Transaction data"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /transactions [post]
func (controller *TransactionsController) Post(c echo.Context) error {
	registerTransactionRequest := new(requests.RegisterTransactionRequest)

	if err := c.Bind(registerTransactionRequest); err != nil {
		return err
	}

	if err := registerTransactionRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if t, err := controller.service.Create(registerTransactionRequest.ToModel()); err != nil {
		c.Logger().Errorf("error creating account, err: %v", err)
		return responses.ErrorResponse(c, http.StatusBadRequest, "Unable to register transaction, try again in a few moments")
	} else {
		return responses.MessageResponse(c, http.StatusCreated, "Transaction with ID:"+t.Id.String()+" has been registered")
	}

}
