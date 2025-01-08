package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// Application struct to describe Application object.
type Application struct {
	models.Model
	ID           uint   `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Name         string `db:"name" json:"name" validate:"required,lte=255"`
	ClientID     string `db:"client_id" json:"client_id" validate:"required,lte=255"`
	ClientSecret string `db:"client_secret" json:"client_secret" validate:"required,lte=255"`
}
