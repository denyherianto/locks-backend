package controllers

import (
	"log"
	"strconv"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
)

// GetUserRole func gets userRole by given ID or 404 error.
// @Description Get userRole by given ID.
// @Summary get userRole by given ID
// @Tags UserRoles
// @Accept json
// @Produce json
// @Param id path string true "UserRole ID"
// @Success 200 {object} entities.UserRole
// @Router /v1/companies/{id} [get]
func GetUserRolesByUserID(c *fiber.Ctx) error {
	// Catch userRole ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Get userRole by ID.
	userRole, err := queries.GetUserRolesByUserID(uint(id))
	if err != nil {
		// Return, if userRole not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "userRole with the given ID is not found",
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": userRole,
		"meta": fiber.Map{},
	})
}

// CreateUserRole func for creates a new userRole.
// @Description Assign a new Role to User.
// @Summary create a new userRole
// @Tags UserRoles
// @Accept json
// @Produce json
// @Param user_id body string true "User ID"
// @Param role_id body string true "Role ID"
// @Success 200 {object} entities.UserRole
// @Security ApiKeyAuth
// @Router /v1/roles/assign [post]
func AssignUserRole(c *fiber.Ctx) error {
	// Create new UserRole struct
	userRole := &entities.UserRole{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&userRole); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a UserRole model.
	validate := utils.NewValidator()

	// Validate userRole fields.
	if err := validate.Struct(userRole); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Get userRole by ID.
	log.Printf("userRole: %+v", userRole)
	foundedUserRoles, _ := queries.GetUserRolesByUserIDAndRoleID(userRole.UserID, userRole.RoleID)
	if len(foundedUserRoles) > 0 {
		// Return, if userRole not found.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": "role already assigned",
		})
	}

	// Create userRole by given model.
	if err := queries.AssignUserRole(userRole); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": userRole,
		"meta": fiber.Map{},
	})
}
