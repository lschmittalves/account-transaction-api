package seeders

import (
	"account-transaction-api/internal/models"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"log"
)

type OperationTypeSeeder struct {
	DB *gorm.DB
}

func NewOperationTypeSeeder(db *gorm.DB) *OperationTypeSeeder {
	return &OperationTypeSeeder{DB: db}
}

func (seeder *OperationTypeSeeder) SetDefaultOperationTypes() {
	operations := map[string]map[string]string{
		"ef4dc378-e57e-4951-ad43-77b8d4af403d": {
			"description": "COMPRA A VISTA",
		},
		"443a4215-80db-4614-888c-dc9be9b29656": {
			"description": "COMPRA PARCELADA",
		},
		"6f025e29-937f-4ca1-af4e-4fa03838f27e": {
			"description": "SAQUE",
		},
		"fce2fa7e-a698-40c8-a765-268d13190329": {
			"description": "PAGAMENTO",
		},
	}

	for key, value := range operations {
		var op = &models.OperationType{}
		r := seeder.DB.Where("id = ? ", key).First(op)
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			op.Id, _ = uuid.Parse(key)
			op.Description = value["description"]

			log.Printf("seeding operation type %v", op)

			seeder.DB.Create(&op)
		}
	}
}
