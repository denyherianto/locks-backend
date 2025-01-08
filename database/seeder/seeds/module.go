package seeds

import (
	"fmt"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

func CreateModuleSeeds() {
 // Create sample roles
 modules := []entities.Module{
  {ID: 1, Name: "Auth", ApplicationID: 2},
  {ID: 2, Name: "Module", ApplicationID: 2},
  {ID: 3, Name: "Settings", ApplicationID: 2},
  {ID: 4, Name: "Master Data", ApplicationID: 3},
 }

 // Create module in the database
 for _, module := range modules {
  err := database.DBManager.Save(&module).Error
  if err != nil {
   fmt.Printf("Error when create module: %s\n", module.Name)
  } else {
   fmt.Printf("Success create module: %s\n", module.Name)
  }
 }
}
