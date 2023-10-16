package account

import (
	"account-transaction-api/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ServiceWrapper interface {
	Create(account *models.Account) (*models.Account, error)
	GetById(id uuid.UUID) (*models.Account, error)
}

type Service struct {
	DB *gorm.DB
}

func NewAccountService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
