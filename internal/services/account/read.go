package account

import (
	"account-transaction-api/internal/models"
	"github.com/google/uuid"
)

func (s *Service) GetById(id uuid.UUID) (*models.Account, error) {
	var acc = &models.Account{}
	err := s.DB.Where("id = ? ", id).Find(acc).Error

	if err != nil {
		acc = nil
	}

	return acc, err
}

func (s *Service) Exists(doc string) bool {
	var acc = &models.Account{}
	err := s.DB.Where("tax_document = ? ", doc).Find(acc).Error
	return err == nil
}
