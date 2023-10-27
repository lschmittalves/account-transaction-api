package account

import (
	"account-transaction-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

func (s *ServiceWrapper) UpdateCreditLimit(id uuid.UUID, transactionValue int64) error {

	acc := models.Account{}

	if err := s.accountReader.GetById(&acc, id); err != nil {
		return err
	}

	if transactionValue < 0 && (transactionValue*-1) > acc.CreditLimit {
		return fmt.Errorf("account id %s does`t have available limit for this transaction", id)
	}

	acc.CreditLimit += transactionValue

	s.accountWriter.Update(&acc)

	return nil
}
