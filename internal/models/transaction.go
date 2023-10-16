package models

import uuid "github.com/satori/go.uuid"

type Transaction struct {
	Base
	Amount          uint64        `json:"amount" gorm:"not null"`
	AccountId       uuid.UUID     `json:"account_id"`
	Account         Account       `json:"-" gorm:"foreignkey:AccountId"`
	OperationTypeId uuid.UUID     `json:"operation_type_id"`
	OperationType   OperationType `json:"-" gorm:"foreignkey:OperationTypeId"`
}
