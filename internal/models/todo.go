package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	Id          string `gorm:"primaryKey"`
	Title       string
	Description string
	Closed      bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time `gorm:"default:null"`
}

func (Todo) TableName() string {
	return "todos"
}

func (t *Todo) BeforeCreate(_ *gorm.DB) error {
	t.Id = uuid.NewString()
	return nil
}
