package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// Role struct to describe Role object.
type Role struct {
	models.Model
	ID             uint              `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Name           string            `db:"name" json:"name" validate:"required,lte=255"`
	Slug           string            `db:"slug" json:"slug" validate:"required,lte=255"`
	RolePermission *[]RolePermission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role_permission"`
}
