package requests

import (
	"account-transaction-api/internal/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type RegisterTransactionRequest struct {
	AccountId       uuid.UUID `json:"account_id" validate:"required" example:"eb8fff89-944c-4da7-b682-cbb1800e36a2"`
	OperationTypeId uuid.UUID `json:"operation_type_id" validate:"required" example:"ef4dc378-e57e-4951-ad43-77b8d4af403d"`
	// amount must be sent in base 100, example: 2.00 is 200
	Amount int64 `json:"amount" validate:"required" example:"20"`
}

func (t RegisterTransactionRequest) Validate() error {
	return validation.ValidateStruct(&t,

		validation.Field(&t.AccountId,
			validation.Required),

		validation.Field(&t.OperationTypeId,
			validation.Required),

		validation.Field(&t.Amount,
			validation.Required,
			validation.Min(1).Error("amount can not be less than 1 cent")),
	)
}

func (t RegisterTransactionRequest) ToModel() *models.Transaction {
	return &models.Transaction{AccountId: t.AccountId, OperationTypeId: t.OperationTypeId, Amount: t.Amount}
}
