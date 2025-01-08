package controllers

import (
	"fmt"
	"time"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateUserRole func for creates a new userKeyCopy.
// @Description Assign a new Role to User.
// @Summary create a new userKeyCopy
// @Tags UserRoles
// @Accept json
// @Produce json
// @Param user_id body string true "User ID"
// @Param role_id body string true "Role ID"
// @Success 200 {object} entities.UserRole
// @Security ApiKeyAuth
// @Router /v1/roles/assign [post]
func GrantUserKeyCopy(c *fiber.Ctx) error {
	// Create new UserKeyCopy struct
	userKeyCopy := &entities.UserKeyCopy{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&userKeyCopy); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a UserRole model.
	validate := utils.NewValidator()

	// Validate userKeyCopy fields.
	if err := validate.Struct(userKeyCopy); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	existing, _ := queries.GetUserKeyCopy(userKeyCopy.UserID, userKeyCopy.KeyCopyID)
	fmt.Printf("existing: %v\n", existing)
	if (existing != entities.UserKeyCopy{}) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "user key copy already exists",
		})
	}

	userKeyCopy.GrantedAt = time.Now()
	userKeyCopy.RevokedAt = nil

	// Get userKeyCopy by ID.
	if userKeyCopy.KeyCopyID != 0 {
		// Get key copy by ID.
		keyCopy, err := queries.GetKeyCopy(userKeyCopy.KeyCopyID)
		if err != nil {
			// Return, if key copy not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   fiber.StatusNotFound,
				"message": "key copies with the given ID is not found",
			})
		}

		userKeyCopy.KeyCopy = &keyCopy
	}

	// Create userKeyCopy by given model.
	if err := queries.CreateUserKeyCopy(userKeyCopy); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"message": "user key copy granted",
	})
}

// CreateUserRole func for creates a new userKeyCopy.
// @Description Assign a new Role to User.
// @Summary create a new userKeyCopy
// @Tags UserRoles
// @Accept json
// @Produce json
// @Param user_id body string true "User ID"
// @Param role_id body string true "Role ID"
// @Success 200 {object} entities.UserRole
// @Security ApiKeyAuth
// @Router /v1/roles/assign [post]
func RevokeUserKeyCopy(c *fiber.Ctx) error {
	// Create new UserKeyCopy struct
	userKeyCopy := &entities.UserKeyCopy{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&userKeyCopy); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a UserRole model.
	validate := utils.NewValidator()

	// Validate userKeyCopy fields.
	if err := validate.Struct(userKeyCopy); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	existing, err := queries.GetUserKeyCopy(userKeyCopy.UserID, userKeyCopy.KeyCopyID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": err.Error(),
		})
	}
	now := time.Now()
	existing.RevokedAt = &now

	if err := queries.UpdateUserKeyCopy(userKeyCopy.UserID, userKeyCopy.KeyCopyID, &existing); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"message": "user key copy revoked",
	})
}
