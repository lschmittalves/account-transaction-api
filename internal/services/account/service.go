package account

import (
	"account-transaction-api/internal/models"
	"account-transaction-api/internal/repositories"
)

type Service interface {
	Create(account *models.Account) (*models.Account, error)
}

type ServiceWrapper struct {
	accountReader repositories.AccountReader
	accountWriter repositories.AccountWriter
}

func NewAccountService(reader repositories.AccountReader, writer repositories.AccountWriter) *ServiceWrapper {
	return &ServiceWrapper{accountReader: reader, accountWriter: writer}
}
