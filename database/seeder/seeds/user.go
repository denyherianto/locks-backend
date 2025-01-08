package seeds

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

func CreateUserSeeds() {
	// Create sample users
	users := []entities.User{
		{ID: 1, Name: "John Doe", Email: "johndoe@denyherianto.com", PasswordHash: utils.GeneratePassword("password"), ResetToken: uuid.New().String(), VerificationToken: uuid.New().String(), UserStatus: 1},
	}

	// Create roles in the database
	for _, user := range users {
		err := database.DBManager.Save(&user).Error
		if err != nil {
			fmt.Printf("Error when create user: %s\n", user.Name)
		} else {
			fmt.Printf("Success create user: %s\n", user.Name)
		}
	}
}
