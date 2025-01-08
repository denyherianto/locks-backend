package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetUserByIDs query for getting one User by given ID.
func GetApplicationByClientID(clientID string) (*entities.Application, error) {
	// Define User variable.
	var application *entities.Application

	// Send query to database.
	if err := database.DBManager.Where("client_id = ?", clientID).First(&application).Error; err != nil {
		// Return empty object and error.
		return nil, err
	}

	return application, nil
}
