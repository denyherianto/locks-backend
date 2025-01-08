package controllers

import (
	"strconv"
	"time"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
)

// GetRoles func gets all exists roles.
// @Description Get all exists roles.
// @Summary get all exists roles
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200 {array} entities.Role
// @Router /v1/roles [get]
func GetRoles(c *fiber.Ctx) error {
	// Get all roles.
	roles, err := queries.GetRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": roles,
		"meta": fiber.Map{
			"total": len(roles),
		},
	})
}

// GetRole func gets role by given ID or 404 error.
// @Description Get role by given ID.
// @Summary get role by given ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} entities.Role
// @Router /v1/roles/{id} [get]
func GetRole(c *fiber.Ctx) error {
	// Catch role ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Get role by ID.
	role, err := queries.GetRole(uint(id))
	if err != nil {
		// Return, if role not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "role with the given ID is not found",
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": role,
		"meta": fiber.Map{},
	})
}

// CreateRole func for creates a new role.
// @Description Create a new role.
// @Summary create a new role
// @Tags Roles
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Success 200 {object} entities.Role
// @Security ApiKeyAuth
// @Router /v1/roles [post]
func CreateRole(c *fiber.Ctx) error {
	// Create new Role struct
	role := &entities.Role{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&role); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a Role model.
	validate := utils.NewValidator()

	// Validate role fields.
	if err := validate.Struct(role); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Create role by given model.
	if err := queries.CreateRole(role); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": role,
		"meta": fiber.Map{},
	})
}

// UpdateRole func for updates role by given ID.
// @Description Update role.
// @Summary update role
// @Tags Roles
// @Accept json
// @Produce json
// @Param id body string true "Role ID"
// @Param name body string true "Name"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/roles/{id} [put]
func UpdateRole(c *fiber.Ctx) error {
	// Catch role ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Create a new user auth struct.
	role := &entities.Role{}

	// Checking received data from JSON body.
	if err := c.BodyParser(role); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Checking, if role with given ID is exists.
	foundedRole, err := queries.GetRole(uint(id))
	if err != nil {
		// Return status 404 and role not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "role with this ID not found",
		})
	}

	// Set initialized data for role:
	foundedRole.Name = role.Name
	foundedRole.UpdatedAt = time.Now()

	// Create a new validator for a Role model.
	validate := utils.NewValidator()
	// Validate role fields.
	if err := validate.Struct(foundedRole); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Update role by given ID.
	if err := queries.UpdateRole(uint(id), foundedRole); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": foundedRole,
		"meta": fiber.Map{},
	})
}

// DeleteRole func for deletes role by given ID.
// @Description Delete role by given ID.
// @Summary delete role by given ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id body string true "Role ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/roles/{id} [delete]
func DeleteRole(c *fiber.Ctx) error {
	// Catch role ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Checking, if role with given ID is exists.
	foundedRole, err := queries.GetRole(uint(id))
	if err != nil {
		// Return status 404 and role not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "role with this ID not found",
		})
	}

	if err := queries.DeleteRole(foundedRole.ID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"message": "Successfully deleted",
	})
}
