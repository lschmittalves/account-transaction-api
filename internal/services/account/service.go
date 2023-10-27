package account

import (
	"account-transaction-api/internal/models"
	"account-transaction-api/internal/repositories"
	"github.com/google/uuid"
)

type Service interface {
	Create(account *models.Account) (*models.Account, error)
	UpdateCreditLimit(id uuid.UUID, transactionValue int64) error
}

type ServiceWrapper struct {
	accountReader repositories.AccountReader
	accountWriter repositories.AccountWriter
}

func NewAccountService(reader repositories.AccountReader, writer repositories.AccountWriter) *ServiceWrapper {
	return &ServiceWrapper{accountReader: reader, accountWriter: writer}
}
