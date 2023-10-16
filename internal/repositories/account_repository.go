package repositories

import (
	"account-transaction-api/internal/models"
	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type AccountRepositoryQ interface {
	GetById(account *models.Account, id uuid.UUID)
	Add(account *models.Account)
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) GetById(a *models.Account, id uuid.UUID) error {
	return r.DB.Where("id = ? ", id).Find(a).Error
}

func (r *AccountRepository) Add(a *models.Account) error {

	if err := r.DB.Create(a).Error; err != nil {
		return err
	}
	if err := r.DB.Save(a).Error; err != nil {
		return err
	}

	return nil
}
