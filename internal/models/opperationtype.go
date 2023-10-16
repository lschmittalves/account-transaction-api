package models

type OperationType struct {
	Base
	Description string `json:"description" gorm:"type:text;not null"`
	IsDebit     bool   `json:"is_debit" gorm:"not null"`
}
