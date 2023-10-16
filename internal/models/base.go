package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Base struct {
	Id        uuid.UUID  `json:"id";gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	if base.Id == uuid.Nil {
		base.Id = uuid.New()
	}
	return scope.SetColumn("Id", base.Id)
}
