package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/models/requests"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetRoles method for getting all roles.
func GetUserRolesByUserID(id uint) ([]requests.UserRoleResponse, error) {
	// Define roles variable.
	var roles []requests.UserRoleResponse
	// Send query to database.
	result := database.DBManager.Model(&entities.UserRole{}).
		Select("user_roles.*,roles.name as role_name").
		Where("user_roles.user_id = ?", id).
		Joins("LEFT JOIN users ON user_roles.user_id = users.id").
		Joins("LEFT JOIN roles ON user_roles.role_id = roles.id").
		Find(&roles)
	if result.Error != nil {
		// Return empty object and error.
		return nil, result.Error
	}
	// Return query result.
	return roles, nil
}

// GetRole method for getting one role by given ID.
func GetUserRolesByUserIDAndRoleID(user_id uint, role_id uint) ([]entities.UserRole, error) {
	// Define roles variable.
	var roles []entities.UserRole
	// Send query to database.
	result := database.DBManager.Where("user_id = ? AND role_id = ?", user_id, role_id).Find(&roles)
	if result.Error != nil {
		// Return empty object and error.
		return nil, result.Error
	}
	// Return query result.
	return roles, nil
}

// AssignUserRole method for creating role by given Role object.
func AssignUserRole(userRole *entities.UserRole) error {
	// Send query to database.
	result := database.DBManager.Create(userRole)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}
