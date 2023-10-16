package account

import (
	"account-transaction-api/internal/models"
	"fmt"
)

func (s *ServiceWrapper) Create(acc *models.Account) (*models.Account, error) {

	if s.accountReader.ExistsDocument(acc.TaxDocument) {
		return nil, fmt.Errorf("account with document %s allready exists", acc.TaxDocument)
	}
	return acc, s.accountWriter.Add(acc)
}
