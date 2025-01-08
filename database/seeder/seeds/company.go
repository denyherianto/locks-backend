package seeds

import (
	"fmt"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

func CreateCompanySeeds() {
 // Create sample companies
 companies := []entities.Company{
  {ID: 1, Name: "Company ABC"},
  {ID: 2, Name: "Company XYZ"},
 }

 // Create companies in the database
 for _, company := range companies {
  err := database.DBManager.Save(&company).Error
  if err != nil {
   fmt.Printf("Error when create company: %s\n", company.Name)
  } else {
   fmt.Printf("Success create company: %s\n", company.Name)
  }
 }
}
