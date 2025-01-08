package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// Company struct to describe Company object.
type Company struct {
	models.Model
	ID   uint   `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Name string `db:"name" json:"name" validate:"required,lte=255"`
}
