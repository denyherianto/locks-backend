package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetKeys method for getting all keys.
func GetKeys(search string, pagination *utils.Pagination) ([]entities.Key, error) {
	// Define keys variable.
	var keys []entities.Key

	// Send query to database.
	result := database.DBManager.Preload("Company").Select("keys.*").Scopes(utils.Paginate(&keys, pagination, database.DBManager)).Where("keys.name LIKE ?", "%"+search+"%").Find(&keys)

	if result.Error != nil {
		// Return empty object and error.
		return nil, result.Error
	}
	// Return query result.
	return keys, nil
}

// GetKey method for getting one key by given ID.
func GetKey(id uint) (entities.Key, error) {
	// Define key variable.
	var key entities.Key

	// Send query to database.
	result := database.DBManager.Preload("Company").First(&key, id)
	if result.Error != nil {
		// Return empty object and error.
		return key, result.Error
	}
	// Return query result.
	return key, nil
}

// CreateKey method for creating key by given Key object.
func CreateKey(key *entities.Key) error {
	// Send query to database.
	result := database.DBManager.Create(key)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// UpdateKey method for updating key by given Key object.
func UpdateKey(id uint, key entities.Key) error {
	// Send query to database.
	result := database.DBManager.Where("id = ?", id).Updates(&key)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// DeleteKey method for delete key by given ID.
func DeleteKey(id uint) error {
	key := entities.Key{}

	// Send query to database.
	result := database.DBManager.Delete(&key, id)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}
