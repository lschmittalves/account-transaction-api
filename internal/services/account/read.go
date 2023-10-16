package account

import (
	"account-transaction-api/internal/models"
	"github.com/google/uuid"
)

func (s *Service) GetById(id uuid.UUID) (acc *models.Account, err error) {
	err = s.DB.Where("id = ? ", id).Find(acc).Error
	return acc, err
}
