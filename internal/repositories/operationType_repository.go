package repositories

import (
	"account-transaction-api/internal/models"
	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type OperationTypeReader interface {
	GetById(op *models.OperationType, id uuid.UUID) error
}

type OperationTypeRepository struct {
	DB *gorm.DB
}

func NewOperationTypeRepository(db *gorm.DB) *OperationTypeRepository {
	return &OperationTypeRepository{DB: db}
}

func (r *OperationTypeRepository) GetById(op *models.OperationType, id uuid.UUID) error {
	return r.DB.Where("id = ? ", id).First(op).Error
}
