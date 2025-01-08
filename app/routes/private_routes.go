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

	// Routes for Keys
	routeV1.Get("/keys", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetKeys)          // get All Keys
	routeV1.Get("/keys/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetKey)       // get Key by ID
	routeV1.Post("/keys", middleware.JWTProtected(), middleware.TokenValidation, controllers.CreateKey)       // create new Key
	routeV1.Put("/keys/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.UpdateKey)    // update Key
	routeV1.Delete("/keys/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.DeleteKey) // delete Key

	// Routes for KeyCopies
	routeV1.Get("/key-copies", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetKeyCopies)              // get All KeyCopies
	routeV1.Get("/key-copies/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.GetKeyCopy)            // get KeyCopy by ID
	routeV1.Post("/key-copies", middleware.JWTProtected(), middleware.TokenValidation, controllers.CreateKeyCopy)            // create new KeyCopy
	routeV1.Post("/key-copies/grant", middleware.JWTProtected(), middleware.TokenValidation, controllers.GrantUserKeyCopy)   // grant UserKeyCopy
	routeV1.Post("/key-copies/revoke", middleware.JWTProtected(), middleware.TokenValidation, controllers.RevokeUserKeyCopy) // revoke UserKeyCopy
	routeV1.Put("/key-copies/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.UpdateKeyCopy)         // update KeyCopy
	routeV1.Delete("/key-copies/:id", middleware.JWTProtected(), middleware.TokenValidation, controllers.DeleteKeyCopy)      // delete KeyCopy
}
