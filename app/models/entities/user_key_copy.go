package entities

import (
	"time"

	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// UserKeyCopy struct to describe User Key Copies object relation.
type UserKeyCopy struct {
	models.Model
	ID        uint       `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	UserID    uint       `db:"user_id" json:"user_id" validate:"required,number"`
	User      *User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	KeyCopyID uint       `db:"key_copy_id" json:"key_copy_id" validate:"required,number"`
	KeyCopy   *KeyCopy   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"key_copy"`
	GrantedAt time.Time  `gorm:"default:current_timestamp" db:"granted_at" json:"granted_at"`
	RevokedAt *time.Time `gorm:"default:current_timestamp" db:"revoked_at" json:"revoked_at"`
}
