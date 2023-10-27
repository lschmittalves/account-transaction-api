package transactions

import (
	"account-transaction-api/internal/models"
	"account-transaction-api/internal/repositories"
	"account-transaction-api/internal/services/account"
)

type Service interface {
	Create(account *models.Transaction) (*models.Transaction, error)
}

type ServiceWrapper struct {
	accountReader     repositories.AccountReader
	accountService    account.Service
	operationReader   repositories.OperationTypeReader
	transactionWriter repositories.TransactionWriter
}

func NewTransactionService(accountReader repositories.AccountReader, accountService account.Service, operationReader repositories.OperationTypeReader, transactionWriter repositories.TransactionWriter) *ServiceWrapper {
	return &ServiceWrapper{accountReader, accountService, operationReader, transactionWriter}
}
