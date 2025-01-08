package seeds

import (
	"fmt"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

func CreateApplicationSeeds() {
 // Create sample companies
 applications := []entities.Application{
  {ID: 1, Name: "Service API", ClientID: "service-api", ClientSecret: utils.GeneratePassword("yQ2kA071Vj2hn73aahUc")},
  {ID: 2, Name: "Main API", ClientID: "eoffice-api", ClientSecret: utils.GeneratePassword("R4DqwhyFtPyYZOiNAVES")},
  {ID: 3, Name: "Secondary API", ClientID: "eproc-api", ClientSecret: utils.GeneratePassword("wErBgoFrdzDd2GRmaGgp")},
 }

 // Create companies in the database
 for _, application := range applications {
  err := database.DBManager.Save(&application).Error
  if err != nil {
   fmt.Printf("Error when create application: %s\n", application.Name)
  } else {
   fmt.Printf("Success create application: %s\n", application.Name)
  }
 }
}
