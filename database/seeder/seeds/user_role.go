package seeds

import (
	"fmt"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

func CreateUserRoleSeeds() {
	// Create sample user roles
	userRoles := []entities.UserRole{
		{UserID: 1, RoleID: 1},
		{UserID: 1, RoleID: 2},
	}

	// Create user role in the database
	for _, userRole := range userRoles {
		err := database.DBManager.Save(&userRole).Error
		if err != nil {
			fmt.Printf("Error when create user role: %v\n", userRole)
		} else {
			fmt.Printf("Success create user role: %v\n", userRole)
		}

	}
}
