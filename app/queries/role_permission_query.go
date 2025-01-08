package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/models/requests"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetRolePermissionByUserID method for getting all role permissions.
func GetRolePermissionByUserID(id uint) ([]requests.UserRolePermissionResponse, error) {
	// Define roles variable.
	var rolePermissions []requests.UserRolePermissionResponse
	// Send query to database.
	result := database.DBManager.Model(&entities.UserRole{}).
		Select("modules.id as module_id, modules.name as module_name, role_permissions.create as permission_create, role_permissions.read as permission_read, role_permissions.update as permission_update, role_permissions.delete as permission_delete").
		// Group("modules").
		Where("user_roles.user_id = ?", id).
		Joins("LEFT JOIN users ON user_roles.user_id = users.id").
		Joins("LEFT JOIN roles ON user_roles.role_id = roles.id").
		Joins("LEFT JOIN role_permissions ON role_permissions.role_id = roles.id").
		Joins("LEFT JOIN modules ON role_permissions.module_id = modules.id").
		Find(&rolePermissions)
	if result.Error != nil {
		// Return empty object and error.
		return nil, result.Error
	}
	// Return query result.
	return rolePermissions, nil
}
