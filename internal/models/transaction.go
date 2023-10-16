package models

import "github.com/google/uuid"

type Transaction struct {
	Base
	Amount          int64         `json:"amount" gorm:"not null"`
	AccountId       uuid.UUID     `json:"account_id"`
	Account         Account       `json:"-" gorm:"foreignkey:AccountId"`
	OperationTypeId uuid.UUID     `json:"operation_type_id"`
	OperationType   OperationType `json:"-" gorm:"foreignkey:OperationTypeId"`
}
