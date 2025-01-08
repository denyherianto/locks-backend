package seeds

import (
	"fmt"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

func CreateRolePermissionSeeds() {
	// Create sample roles
	rolePermissions := []entities.RolePermission{
		{RoleID: 1, ModuleID: 1, Create: true, Read: true, Update: true, Delete: true},
		{RoleID: 1, ModuleID: 2, Create: true, Read: true, Update: false, Delete: false},
		{RoleID: 2, ModuleID: 1, Create: true, Read: false, Update: false, Delete: false},
		{RoleID: 2, ModuleID: 2, Create: true, Read: false, Update: false, Delete: false},
	}

	// Create role permission in the database
	for _, rolePermission := range rolePermissions {
		err := database.DBManager.Save(&rolePermission).Error
		if err != nil {
			fmt.Printf("Error when create role permission: %v\n", rolePermission)
		} else {
			fmt.Printf("Success create role permission: %v\n", rolePermission)
		}
	}
}
