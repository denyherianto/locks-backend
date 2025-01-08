package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// User struct to describe User object.
type User struct {
	models.Model
	ID                uint        `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	Name              string      `db:"name" json:"name" validate:"required,lte=255"`
	Username          string      `gorm:"unique" db:"username" json:"username" validate:"required,gte=3,lte=20"`
	Email             string      `gorm:"index" db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash      string      `db:"password_hash" json:"-" validate:"required,lte=255"`
	ResetToken        string      `db:"reset_token" json:"-" validate:"required,lte=255"`
	VerificationToken string      `db:"verification_token" json:"-" validate:"required,lte=255"`
	UserStatus        int         `db:"user_status" json:"user_status" validate:"required,len=1"`
	Roles             *[]UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"roles"`
}
