package routes

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/controllers"
	"github.com/denyherianto/go-fiber-boilerplate/app/middleware"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// API V1 Routes Group
	routeV1 := a.Group("/api/v1")

	// Routes for Auth
	routeV1.Post("/auth/logout", middleware.JWTProtected(), middleware.TokenValidation, controllers.UserSignOut) // de-authorization user

	// Routes for Token
	routeV1.Post("/token/renew", middleware.JWTProtected(), middleware.TokenValidation, controllers.RenewTokens)  // renew Access & Refresh tokens
	routeV1.Post("/token/verify", middleware.JWTProtected(), middleware.TokenValidation, controllers.VerifyToken) // verify Access token

	// Routes for Companies
	routeV1.Get("/companies", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetCompanies)         // get All Companies
	routeV1.Get("/companies/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetCompany)       // get Company by ID
	routeV1.Post("/companies", middleware.JWTProtected(), middleware.TokenValidation, controllers.CreateCompany)       // create new Company
	routeV1.Put("/companies/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.UpdateCompany)    // update Company
	routeV1.Delete("/companies/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.DeleteCompany) // delete Company

	// Routes for ErrorLogs
	routeV1.Get("/error-logs", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetErrorLogs)          // get All error logs
	routeV1.Get("/error-logs/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetErrorLog)       // get ErrorLog by ID
	routeV1.Post("/error-logs", middleware.JWTProtected(), middleware.TokenValidation, controllers.CreateErrorLog)       // create new ErrorLog
	routeV1.Delete("/error-logs/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.DeleteErrorLog) // delete ErrorLog

	// Routes for Roles
	routeV1.Get("/roles", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetRoles)               // get All Roles
	routeV1.Get("/roles/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetRole)            // get Role by ID
	routeV1.Post("/roles", middleware.JWTProtected(), middleware.TokenValidation, controllers.CreateRole)            // create new Role
	routeV1.Put("/roles/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.UpdateRole)         // update Role
	routeV1.Delete("/roles/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.DeleteRole)      // delete Role
	routeV1.Post("/roles/assign", middleware.JWTProtected(), middleware.TokenValidation, controllers.AssignUserRole) // de-authorization user
}
