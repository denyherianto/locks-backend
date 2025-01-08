package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetKeyCopies method for getting all keyCopies.
func GetKeyCopies(search string, pagination *utils.Pagination) ([]entities.KeyCopy, error) {
	// Define keyCopies variable.
	var keyCopies []entities.KeyCopy

	// Send query to database.
	result := database.DBManager.Preload("Key").Select("key_copies.*").Scopes(utils.Paginate(&keyCopies, pagination, database.DBManager)).Where("key_copies.identifier LIKE ?", "%"+search+"%").Find(&keyCopies)

	if result.Error != nil {
		// Return empty object and error.
		return nil, result.Error
	}
	// Return query result.
	return keyCopies, nil
}

// GetKeyCopy method for getting one keyCopy by given ID.
func GetKeyCopy(id uint) (entities.KeyCopy, error) {
	// Define keyCopy variable.
	var keyCopy entities.KeyCopy

	// Send query to database.
	result := database.DBManager.Preload("Key").First(&keyCopy, id)
	if result.Error != nil {
		// Return empty object and error.
		return keyCopy, result.Error
	}
	// Return query result.
	return keyCopy, nil
}

// CreateKeyCopy method for creating keyCopy by given KeyCopy object.
func CreateKeyCopy(keyCopy *entities.KeyCopy) error {
	// Send query to database.
	result := database.DBManager.Create(keyCopy)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// UpdateKeyCopy method for updating keyCopy by given KeyCopy object.
func UpdateKeyCopy(id uint, keyCopy entities.KeyCopy) error {
	// Send query to database.
	result := database.DBManager.Where("id = ?", id).Updates(&keyCopy)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// DeleteKeyCopy method for delete keyCopy by given ID.
func DeleteKeyCopy(id uint) error {
	keyCopy := entities.KeyCopy{}

	// Send query to database.
	result := database.DBManager.Delete(&keyCopy, id)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}
