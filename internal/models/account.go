package models

type Account struct {
	Base
	Name        string `json:"name" gorm:"type:varchar(200);not null"`
	TaxDocument string `json:"tax_document" gorm:"uniqueIndex;type:varchar(50);"`
	CreditLimit int64  `json:"amount" gorm:"not null;default:0"`
}
