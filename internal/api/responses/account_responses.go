package responses

import (
	"account-transaction-api/internal/models"
	"github.com/google/uuid"
	"time"
)

type AccountResponse struct {
	AccountId      uuid.UUID `json:"account_id" example:"eb8fff89-944c-4da7-b682-cbb1800e36a2"`
	DocumentNumber string    `json:"document_number" example:"99995555444"`
	CreatedAt      time.Time `json:"created_at" example:"2023-10-16T04:57:33.299641Z"`
	UpdatedAt      time.Time `json:"updated_at" example:"2023-10-16T04:57:33.299641Z"`
	Name           string    `json:"name" example:"John Doe"`
}

func NewAccountResponse(account *models.Account) *AccountResponse {
	return &AccountResponse{
		AccountId:      account.Id,
		Name:           account.Name,
		DocumentNumber: account.TaxDocument,
		CreatedAt:      account.CreatedAt,
		UpdatedAt:      account.UpdatedAt}
}
