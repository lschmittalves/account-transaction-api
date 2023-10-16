package controllers

import (
	"account-transaction-api/internal/api/requests"
	"account-transaction-api/internal/api/responses"
	"account-transaction-api/internal/models"
	"account-transaction-api/internal/repositories"
	accountService "account-transaction-api/internal/services/account"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AccountsController struct {
	echo    *echo.Echo
	service accountService.Service
	reader  repositories.AccountReader
}

func NewAccountsController(echo *echo.Echo, service accountService.Service, reader repositories.AccountReader) *AccountsController {
	return &AccountsController{echo: echo, service: service, reader: reader}
}

// Post godoc
// @Summary Create a new account
// @Description Create a new account
// @ID account-create
// @Tags Account Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateAccountRequest true "Account registration data"
// @Success 201 {object} responses.AccountResponse
// @Failure 400 {object} responses.Error
// @Router /accounts [post]
func (controller *AccountsController) Post(c echo.Context) error {
	createAccountRequest := new(requests.CreateAccountRequest)

	if err := c.Bind(createAccountRequest); err != nil {
		return err
	}

	if err := createAccountRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if acc, err := controller.service.Create(createAccountRequest.ToModel()); err != nil {
		c.Logger().Errorf("error creating account, err: %v", err)
		return responses.ErrorResponse(c, http.StatusBadRequest, "Unable to create account, try again in a few moments")
	} else {
		return responses.Response(c, http.StatusCreated, responses.NewAccountResponse(acc))
	}

}

// Get godoc
// @Summary Get account by id
// @Description Get one account by id
// @ID account-get
// @Tags Account Actions
// @Produce json
// @Param id path string true "Account UUID"
// @Success 200 {object} responses.AccountResponse
// @NoContent 204
// @Failure 400 {object} responses.Error
// @Router /accounts/{id} [get]
func (controller *AccountsController) Get(c echo.Context) error {

	accountId := c.Param("id")
	accountUuid, err := uuid.Parse(accountId)

	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid account id")
	}

	var acc = &models.Account{}
	if err := controller.reader.GetById(acc, accountUuid); err != nil {
		c.Logger().Errorf("error retrieving account %s, err: %v", accountUuid, err)
		return responses.Response(c, http.StatusNoContent, nil)
	}

	return responses.Response(c, http.StatusOK, responses.NewAccountResponse(acc))
}
