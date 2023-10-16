package account

import "account-transaction-api/internal/models"

func (s *Service) Create(acc *models.Account) (*models.Account, error) {
	return acc, s.DB.Create(acc).Error
}
