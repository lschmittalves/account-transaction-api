package transactions

import (
	"account-transaction-api/internal/models"
	"account-transaction-api/internal/repositories"
)

type Service interface {
	Create(account *models.Transaction) (*models.Transaction, error)
}

type ServiceWrapper struct {
	accountReader     repositories.AccountReader
	operationReader   repositories.OperationTypeReader
	transactionWriter repositories.TransactionWriter
}

func NewTransactionService(accountReader repositories.AccountReader, operationReader repositories.OperationTypeReader, transactionWriter repositories.TransactionWriter) *ServiceWrapper {
	return &ServiceWrapper{accountReader, operationReader, transactionWriter}
}
