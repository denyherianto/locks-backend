package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
)

func PermissionValidation(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Get User Roles by User ID.
	userRoles, _ := queries.GetUserRolesByUserID(claims.UserID)
	log.Printf("User roles: %v", userRoles)

	return c.Next()
}
