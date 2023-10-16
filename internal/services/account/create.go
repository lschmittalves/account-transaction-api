package account

import (
	"account-transaction-api/internal/models"
	"fmt"
)

func (s *Service) Create(acc *models.Account) (*models.Account, error) {

	if s.Exists(acc.TaxDocument) {
		return nil, fmt.Errorf("account with document %s allready exists", acc.TaxDocument)
	}

	return acc, s.DB.Create(acc).Error
}
