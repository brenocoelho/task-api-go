package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key"`
	Username  string     `json:"username" gorm:"type:varchar(255);unique_index;not null"`
	Email     string     `json:"email" gorm:"type:varchar(200);unique_index"`
	Phone     string     `json:"phone" gorm:"type:varchar(30);unique_index"`
	Document  string     `json:"document" gorm:"type:varchar(255);unique_index"`
	FirstName string     `json:"first_name" gorm:"type:varchar(100)"`
	LastName  string     `json:"last_name" gorm:"type:varchar(255)"`
	Password  string     `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
