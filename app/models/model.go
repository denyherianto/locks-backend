package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time      `gorm:"default:current_timestamp" db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp" db:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" db:"deleted_at" json:"-"`
}
