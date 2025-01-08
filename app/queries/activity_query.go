package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// CreateActivity method for creating activity by given Activity object.
func CreateActivity(activity *entities.Activity) error {
	// Send query to database.
	result := database.DBManager.Create(activity)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}
