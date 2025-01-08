package entities

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models"
)

// Activity struct to describe User Activity / Audit Log object.
type Activity struct {
	models.Model
	ID            uint         `gorm:"primaryKey;autoIncrement:true;unique" db:"id" json:"id"`
	UserID        uint         `db:"user_id" json:"user_id" validate:"number"`
	User          *User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	ApplicationID uint         `db:"application_id" json:"-" validate:"number"`
	Application   *Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"application"`
	ModuleName    string       `db:"module_name" json:"module_name" validate:"required,lte=255"`
	FunctionName  string       `db:"function_name" json:"function_name" validate:"lte=255"`
	TableName     string       `db:"table_name" json:"table_name" validate:"lte=255"`
	ReferenceID   uint         `db:"reference_id" json:"reference_id" validate:"number"`
	Description   string       `db:"description" json:"description" validate:"lte=255"`
	OperationType string       `db:"operation_type" json:"operation_type" validate:"lte=255"` // CREATE, READ, UPDATE, DELETE
	BeforeChange  string       `db:"before_change" json:"before_change" gorm:"type:text"`
	AfterChange   string       `db:"after_change" json:"after_change" gorm:"type:text"`
	EndpointURL   string       `db:"endpoint_url" json:"endpoint_url" validate:"lte=255"`
	IPAddress     string       `db:"ip_address" json:"ip_address" validate:"lte=255"`
	UserAgent     string       `db:"user_agent" json:"user_agent" validate:"lte=255"`
}
