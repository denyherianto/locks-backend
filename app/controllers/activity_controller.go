package controllers

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils/logger"
	"github.com/gofiber/fiber/v2"
)

// CreateActivity func for creates a new activity.
// @Description Create a new activity.
// @Summary create a new activity
// @Tags Activities
// @Accept json
// @Produce json
// @Param user_id body string true "User ID"
// @Param application_id body string true "Application ID"
// @Param module_name body string true "Module Name"
// @Param function_name body string true "Function Name"
// @Param table_name body string true "Table Name"
// @Param reference_id body string true "Reference ID"
// @Param operation_type body string true "Operation Type"
// @Param endpoint_url body string true "Endpoint URL"
// @Success 200 {object} entities.Activity
// @Security ApiKeyAuth
// @Router /v1/activities [post]
func CreateActivity(c *fiber.Ctx) error {
	// Create new Activity struct
	activity := &entities.Activity{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&activity); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a Activity model.
	validate := utils.NewValidator()

	// Validate activity fields.
	if err := validate.Struct(activity); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Log Activity
	logger.LogActivity(c, activity)

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": activity,
		"meta": fiber.Map{},
	})
}
