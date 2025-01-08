package entities

import "github.com/denyherianto/go-fiber-boilerplate/app/models"

// User Role struct to describe User and Role relations.
type UserRole struct {
	models.Model
	ID     uint  `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	UserID uint  `db:"user_id" json:"-" validate:"required,number"`
	RoleID uint  `db:"role_id" json:"-" validate:"required,number"`
}
