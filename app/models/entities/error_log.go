package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// Log struct to describe Log object.
type ErrorLog struct {
	models.Model
	ID            uint   `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Level         string `db:"level" json:"level" validate:"required,lte=50"`
	ApplicationID uint   `db:"application_id" json:"application_id" validate:"required,number"`
	ModuleID      uint   `db:"module_id" json:"module_id" validate:"required,number"`
	ActivityID    uint   `db:"activity_id" json:"activity_id" validate:"required,number"`
	UserID        uint   `db:"user_id" json:"user_id" validate:"required,number"`
	CompanyID     uint   `db:"company_id" json:"company_id" validate:"required,number"`
	Description   string `db:"description" json:"description" validate:"lte=255"`
}
