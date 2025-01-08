package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetUserKeyCopy method for getting one keyCopy by given ID.
func GetUserKeyCopy(userId uint, keyCopyId uint) (entities.UserKeyCopy, error) {
	// Define keyCopy variable.
	var userKeyCopy entities.UserKeyCopy

	// Send query to database.
	result := database.DBManager.Preload("User").Preload("KeyCopy").Where("user_id = ? AND key_copy_id = ? AND revoked_at IS NULL", userId, keyCopyId).First(&userKeyCopy)
	if result.Error != nil {
		// Return empty object and error.
		return userKeyCopy, result.Error
	}
	// Return query result.
	return userKeyCopy, nil
}

// AssignUserKeyCopy method for creating user relation with key copy by given UserKeyCopy object.
func CreateUserKeyCopy(userKeyCopy *entities.UserKeyCopy) error {
	// Send query to database.
	result := database.DBManager.Create(userKeyCopy)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// AssignUserKeyCopy method for creating user relation with key copy by given UserKeyCopy object.
func UpdateUserKeyCopy(userId uint, keyCopyId uint, userKeyCopy *entities.UserKeyCopy) error {
	// Send query to database.
	result := database.DBManager.Where("user_id = ? AND key_copy_id = ?", userId, keyCopyId).Updates(&userKeyCopy)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}
