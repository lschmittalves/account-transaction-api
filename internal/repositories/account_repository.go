package repositories

import (
	"account-transaction-api/internal/models"
	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
)

type AccountWriter interface {
	Add(account *models.Account) error
	Update(account *models.Account) error
}

type AccountReader interface {
	GetById(account *models.Account, id uuid.UUID) error
	ExistsDocument(document string) bool
	ExistsId(id uuid.UUID) bool
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) GetById(a *models.Account, id uuid.UUID) error {
	return r.DB.Where("id = ? ", id).First(a).Error
}

func (r *AccountRepository) ExistsDocument(document string) bool {
	var acc = &models.Account{}
	return r.DB.Where("tax_document = ? ", document).First(acc).Error == nil
}

func (r *AccountRepository) ExistsId(id uuid.UUID) bool {
	var acc = &models.Account{}
	return r.DB.Where("id = ? ", id).First(acc).Error == nil
}

func (r *AccountRepository) Add(a *models.Account) error {

	if err := r.DB.Create(a).Error; err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) Update(account *models.Account) error {

	r.DB.Save(&account)
	return nil
}
