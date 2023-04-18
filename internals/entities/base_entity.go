package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID string `gorm:"type:char(36);primary_key" json:"id"`
	// add more common columns like CreatedAt
	CreatedAt *time.Time
	UpdatedAt *time.Time
	// ...
}

// This functions are called before creating Base

func (e *Base) BeforeCreate(db *gorm.DB) (err error) {
	now := time.Now()
	e.ID = uuid.New().String()
	e.CreatedAt = &now
	e.UpdatedAt = &now
	return nil
}

func (e *Base) BeforeSave(db *gorm.DB) (err error) {
	now := time.Now()
	e.UpdatedAt = &now
	return nil
}
