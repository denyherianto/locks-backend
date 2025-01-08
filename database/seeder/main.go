package main

import (
	"fmt"

	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
	"github.com/denyherianto/go-fiber-boilerplate/database/seeder/seeds"
)

func main() {
	database.OpenDBConnection()

	fmt.Printf("\nCreating Applications: \n")
	seeds.CreateApplicationSeeds()

	fmt.Printf("\nCreating Modules: \n")
	seeds.CreateModuleSeeds()

	fmt.Printf("\nCreating Companies: \n")
	seeds.CreateCompanySeeds()

	fmt.Printf("\nCreating Roles: \n")
	seeds.CreateRoleSeeds()

	fmt.Printf("\nCreating Role Permissions: \n")
	seeds.CreateRolePermissionSeeds()

	fmt.Printf("\nCreating Users: \n")
	seeds.CreateUserSeeds()

	fmt.Printf("\nCreating User Roles: \n")
	seeds.CreateUserRoleSeeds()
}
