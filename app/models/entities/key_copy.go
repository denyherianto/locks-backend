package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// Key struct to describe Key object.
type KeyCopy struct {
	models.Model
	ID         uint   `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Identifier string `db:"identifier" json:"identifier" validate:"required,lte=255"`
	Status     string `db:"status" json:"status" validate:"required,lte=10"`
	KeyID      uint   `db:"key_id" json:"key_id" validate:"required,number"`
	Key        *Key   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"key"`
}
