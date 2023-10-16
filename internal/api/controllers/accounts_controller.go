package controllers

import (
	"account-transaction-api/internal/api/requests"
	"account-transaction-api/internal/api/responses"
	accountService "account-transaction-api/internal/services/account"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type AccountsController struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewAccountsController(echo *echo.Echo, db *gorm.DB) *AccountsController {
	return &AccountsController{Echo: echo, DB: db}
}

// Post godoc
// @Summary Create a new account
// @Description Create a new account
// @ID account-create
// @Tags Account Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateAccountRequest
// @Success 201 {object} responses.AccountResponse
// @Failure 400 {object} responses.Error
// @Router /posts [post]
func (controller *AccountsController) Post(c echo.Context) error {
	createAccountRequest := new(requests.CreateAccountRequest)

	if err := c.Bind(createAccountRequest); err != nil {
		return err
	}

	if err := createAccountRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	service := accountService.NewAccountService(controller.DB)

	if acc, err := service.Create(createAccountRequest.ToModel()); err != nil {
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
// @Success 200 {object} responses.AccountResponse
// @Failure 400 {object} responses.Error
// @Router /posts [get]
func (controller *AccountsController) Get(c echo.Context) error {

	accountId := c.Param("id")
	accountUuid := uuid.FromStringOrNil(accountId)

	if accountUuid == uuid.Nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid account id")
	}

	service := accountService.NewAccountService(controller.DB)

	if acc, err := service.GetById(accountUuid); err != nil {
		c.Logger().Errorf("error retrieving account %s, err: %v", accountUuid, err)
		return responses.Response(c, http.StatusNoContent, nil)
	} else if acc != nil {
		return responses.Response(c, http.StatusOK, acc)
	}

	return responses.Response(c, http.StatusNoContent, nil)
}
