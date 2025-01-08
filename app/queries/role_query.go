package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetRoles method for getting all roles.
func GetRoles() ([]entities.Role, error) {
	// Define roles variable.
	var roles []entities.Role
	// Send query to database.
	result := database.DBManager.Find(&roles)
	if result.Error != nil {
		// Return empty object and error.
		return nil, result.Error
	}
	// Return query result.
	return roles, nil
}

// GetRole method for getting one role by given ID.
func GetRole(id uint) (entities.Role, error) {
	// Define role variable.
	var role entities.Role
	// Send query to database.
	result := database.DBManager.First(&role, id)
	if result.Error != nil {
		// Return empty object and error.
		return role, result.Error
	}
	// Return query result.
	return role, nil
}

// CreateRole method for creating role by given Role object.
func CreateRole(role *entities.Role) error {
	// Send query to database.
	result := database.DBManager.Create(role)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// UpdateRole method for updating role by given Role object.
func UpdateRole(id uint, role entities.Role) error {
	// Send query to database.
	result := database.DBManager.Where("id = ?", id).Updates(&role)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// DeleteRole method for delete role by given ID.
func DeleteRole(id uint) error {
	role := entities.Role{}

	// Send query to database.
	result := database.DBManager.Delete(&role, id)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}
