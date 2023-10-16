package models

type Account struct {
	Base
	Name        string `json:"name" gorm:"type:varchar(200);not null"`
	TaxDocument string `json:"tax_document" gorm:"uniqueIndex;type:varchar(50);"`
}
