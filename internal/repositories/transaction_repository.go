package repositories

import (
	"account-transaction-api/internal/models"
	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type TransactionRepositoryQ interface {
	GetById(account *models.Transaction, id uuid.UUID)
	Add(account *models.Transaction)
}

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) GetById(t *models.Transaction, id uuid.UUID) error {
	return r.DB.Where("id = ? ", id).Find(t).Error
}

func (r *TransactionRepository) Add(t *models.Transaction) error {

	if err := r.DB.Create(t).Error; err != nil {
		return err
	}
	if err := r.DB.Save(t).Error; err != nil {
		return err
	}

	return nil
}
