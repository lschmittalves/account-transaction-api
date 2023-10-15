package seeders

import (
	"account-transaction-api/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
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
		op := models.OperationType{}
		r := seeder.DB.First(&op, key)

		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			op.ID = uuid.FromStringOrNil(key)
			op.Description = value["description"]

			fmt.Printf("seeding operation type %v", op)

			seeder.DB.Create(&op)
		}
	}
}
