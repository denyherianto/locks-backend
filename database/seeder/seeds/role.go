package seeds

import (
	"fmt"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

func CreateRoleSeeds() {
 // Create sample roles
 roles := []entities.Role{
  {ID: 1, Name: "Super Admin", Slug: "super-admin"},
  {ID: 2, Name: "Editor", Slug: "editor"},
  {ID: 3, Name: "Contributor", Slug: "contributor"},
 }

 // Create roles in the database
 for _, role := range roles {
  err := database.DBManager.Save(&role).Error
  if err != nil {
   fmt.Printf("Error when create roles: %s\n", role.Name)
  } else {
   fmt.Printf("Success create roles: %s\n", role.Name)
  }

 }
}
