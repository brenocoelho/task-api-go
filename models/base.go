package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
