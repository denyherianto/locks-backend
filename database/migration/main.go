package main

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

func main() {
	database.OpenDBConnection()

	database.DBManager.AutoMigrate(
		&entities.Application{},
		&entities.Module{},
		&entities.Role{},
		&entities.RolePermission{},
		&entities.Company{},
		&entities.ErrorLog{},
		&entities.User{},
		&entities.UserRole{},
		&entities.Activity{},
		&entities.ErrorLog{},
		&entities.Key{},
		&entities.KeyCopy{},
		&entities.UserKeyCopy{},
	)
}
