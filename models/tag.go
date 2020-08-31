package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	UserID    uuid.UUID  `json:"-" gorm:"type:uuid;primary_key"`
	ID        string     `json:"id" gorm:"type:varchar(36);primary_key"`
	Name      string     `json:"name" gorm:"type:varchar(100);not null"`
	Color     string     `json:"color" gorm:"type:varchar(30)"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}
