package requests

import (
	"account-transaction-api/internal/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" validate:"required" example:"123456789"`
	Name           string `json:"name" validate:"required" example:"John Doe"`
}

func (bp CreateAccountRequest) Validate() error {
	return validation.ValidateStruct(&bp,

		validation.Field(&bp.DocumentNumber,
			validation.Required,
			validation.Length(5, 50),
			validation.Match(regexp.MustCompile("^[0-9]*$")).Error("document only accept numbers")),

		validation.Field(&bp.Name,
			validation.Required,
			validation.Length(5, 200)),
	)
}

func (bp CreateAccountRequest) ToModel() *models.Account {
	return &models.Account{TaxDocument: bp.DocumentNumber, Name: bp.Name}
}
