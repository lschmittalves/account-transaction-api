package repositories

import (
	"account-transaction-api/internal/cache"
	"account-transaction-api/internal/models"
	"context"
	"fmt"
	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type TransactionReader interface {
	GetById(transaction *models.Transaction, id uuid.UUID)
}

type TransactionWriter interface {
	Add(t *models.Transaction) error
}

type TransactionRepository struct {
	DB    *gorm.DB
	cache cache.Cache
}

func NewTransactionRepository(db *gorm.DB, cache cache.Cache) *TransactionRepository {
	return &TransactionRepository{DB: db, cache: cache}
}

func (r *TransactionRepository) GetById(t *models.Transaction, id uuid.UUID) error {
	return r.DB.Where("id = ? ", id).Find(t).Error
}

func (r *TransactionRepository) Add(t *models.Transaction) error {

	ctx := context.Background()
	cacheId := fmt.Sprintf("%s/%s/%d", t.AccountId, t.OperationTypeId, t.Amount)

	if !r.cache.SetEx(ctx, cache.TransactionLockKeyPattern, cacheId) {
		return fmt.Errorf("transaction %s has been sent allready", cacheId)
	}

	if err := r.DB.Create(t).Error; err != nil {
		return err
	}
	return nil
}
