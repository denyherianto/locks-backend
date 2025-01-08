package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// RolePermission struct to describe RolePermission object.
type RolePermission struct {
	models.Model
	ID       uint   `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	RoleID   uint   `db:"role_id" json:"-" validate:"required,number"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	ModuleID uint   `db:"module_id" json:"-" validate:"required,number"`
	Module   Module `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"module"`
	Create   bool   `db:"create" json:"create" validate:"required,boolean"`
	Read     bool   `db:"read" json:"read" validate:"required,boolean"`
	Update   bool   `db:"update" json:"update" validate:"required,boolean"`
	Delete   bool   `db:"delete" json:"delete" validate:"required,boolean"`
}
