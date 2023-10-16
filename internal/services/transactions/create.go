package transactions

import (
	"account-transaction-api/internal/models"
	"fmt"
)

func (s *ServiceWrapper) Create(t *models.Transaction) (*models.Transaction, error) {

	if !s.accountReader.ExistsId(t.AccountId) {
		return nil, fmt.Errorf("account with id %s does not exists", t.AccountId)
	}

	var op = &models.OperationType{}
	if err := s.operationReader.GetById(op, t.OperationTypeId); err != nil {
		return nil, fmt.Errorf("operation type with id %s does not exists", t.OperationTypeId)
	}

	if op.IsDebit {
		t.Amount = t.Amount * -1
	}

	return t, s.transactionWriter.Add(t)

}
