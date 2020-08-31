package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Task struct {
	UserID    uuid.UUID      `json:"-" gorm:"type:uuid;primary_key"`
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
	Notes     string         `json:"notes" gorm:"type:text"`
	DueDate   string         `json:"due_date" gorm:"type:varchar(255)"`
	Frequency string         `json:"frequency" gorm:"type:varchar(255)"`
	Status    string         `json:"status" gorm:"type:varchar(10)"`
	Tags      postgres.Jsonb `json:"tags"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `json:"-" sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (task *Task) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
