package responses

import (
	"account-transaction-api/internal/models"
	uuid "github.com/satori/go.uuid"
)

type AccountResponse struct {
	AccountId      uuid.UUID `json:"account_id" example:"eb8fff89-944c-4da7-b682-cbb1800e36a2"`
	DocumentNumber string    `json:"document_number" example:"99995555444"`
	Name           string    `json:"name" example:"John Doe"`
}

func NewAccountResponse(account *models.Account) *AccountResponse {
	return &AccountResponse{AccountId: account.ID, Name: account.Name, DocumentNumber: account.TaxDocument}
}
