package repositories

import (
	"account-transaction-api/internal/models"
	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
)

type OperationTypeRepositoryQ interface {
	GetById(account *models.Account, id uuid.UUID)
}

type OperationTypeRepository struct {
	DB *gorm.DB
}

func NewOperationTypeRepository(db *gorm.DB) *OperationTypeRepository {
	return &OperationTypeRepository{DB: db}
}

func (r *OperationTypeRepository) GetById(opType *models.OperationType, id uuid.UUID) error {
	return r.DB.Where("id = ? ", id).Find(opType).Error
}
