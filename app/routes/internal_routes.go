package routes

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/controllers"
	"github.com/denyherianto/go-fiber-boilerplate/app/middleware"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func InternalRoutes(a *fiber.App) {
	// API V1 Routes Group
	routeV1 := a.Group("/api/v1/internal")

	// Routes for Users
	routeV1.Get("/users", middleware.BasicAuthValidation, controllers.GetUsersByIDs) // Get Users by IDs

	// Routes for Activities
	routeV1.Post("/activities", middleware.BasicAuthValidation, controllers.CreateActivity) // create Activity Log
}
