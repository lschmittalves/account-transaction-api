package models

type OperationType struct {
	Base
	Description string `json:"description" gorm:"type:text;not null"`
}
