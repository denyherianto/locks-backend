package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// Key struct to describe Key object.
type Key struct {
	models.Model
	ID          uint     `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Name        string   `db:"name" json:"name" validate:"required,lte=255"`
	Description string   `db:"name" json:"description"`
	CompanyID   uint     `db:"company_id" json:"company_id" validate:"required,number"`
	Company     *Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company"`
}
