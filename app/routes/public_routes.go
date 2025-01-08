package routes

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Health Check
	a.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// API V1 Routes Group
	routeV1 := a.Group("/api/v1")

	// Routes for Users
	routeV1.Post("/user/register", controllers.UserSignUp) // register a new user
	routeV1.Post("/user/login", controllers.UserSignIn)    // auth, return Access & Refresh tokens
}
