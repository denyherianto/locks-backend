package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetUserByID query for getting one User by given ID.
func GetUserByID(id uint) (*entities.User, error) {
	// Define User variable.
	var user *entities.User

	// Send query to database.
	if err := database.DBManager.First(&user, id).Error; err != nil {
		// Return empty object and error.
		return user, err
	}

	return user, nil
}

// GetUserByIDs query for getting one User by given ID.
func GetUserByIDs(ids []uint) (*[]entities.User, error) {
	// Define User variable.
	var users *[]entities.User

	// Send query to database.
	if err := database.DBManager.Find(&users, ids).Error; err != nil {
		// Return empty object and error.
		return nil, err
	}

	return users, nil
}

// GetUserByEmailOrUsername query for getting one User by given Email or Username.
func GetUserByEmailOrUsername(identifier string) (*entities.User, error) {
	// Define User variable.
	var user *entities.User

	// Send query to database.
	if err := database.DBManager.Where("email = ?", identifier).Or("username = ?", identifier).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser query for creating a new user by given email and password hash.
func CreateUser(user *entities.User) error {
	if err := database.DBManager.Create(&user).Error; err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
