package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// Module struct to describe Module object.
type Module struct {
	models.Model
	ID            uint        `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Name          string      `db:"name" json:"name" validate:"required,lte=255"`
	ApplicationID uint        `db:"application_id" json:"-" validate:"required,number"`
	Application   *Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"application"`
}
