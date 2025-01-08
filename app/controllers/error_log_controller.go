package controllers

import (
	"strconv"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
)

// GetErrorLogs func gets all exists error logs.
// @Description Get all exists error logs.
// @Summary get all exists error logs
// @Tags ErrorLog
// @Accept json
// @Produce json
// @Param search query string false "Search query"
// @Param limit query string false "Limit query"
// @Param offset query string false "Offset query"
// @Success 200 {object} requests.ErrorLogResponse "Success response"
// @Router /v1/error-logs [get]
func GetErrorLogs(c *fiber.Ctx) error {
	search := c.Query("search", "")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	// Create pagination object.
	pagination := &utils.Pagination{
		Page:  page,
		Limit: limit,
	}

	// Get all errorlogs.
	errorlogs, err := queries.GetErrorLogs(search, pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": errorlogs,
		"meta": fiber.Map{
			"current_page": pagination.TotalRows,
			"last_page":    pagination.TotalPages,
			"per_page":     pagination.Page,
			"total":        pagination.Limit,
		},
	})
}

// GetErrorLog func gets error log by given ID or 404 error.
// @Description Get error log by given ID.
// @Summary get error log by given ID
// @Tags ErrorLog
// @Accept json
// @Produce json
// @Param id path int true "ErrorLog ID"
// @Success 200 {object} entities.ErrorLog
// @Router /v1/error-logs/{id} [get]
func GetErrorLog(c *fiber.Ctx) error {
	// Catch error log ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Get errorlog by ID.
	errorlog, err := queries.GetErrorLog(uint(id))
	if err != nil {
		// Return, if error log not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "error log with the given ID is not found",
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": errorlog,
		"meta": fiber.Map{},
	})
}

// CreateErrorLog func for creates a new error log.
// @Description Create a new error log.
// @Summary create a new error log
// @Tags ErrorLog
// @Accept json
// @Produce json
// @Param level body string true "Level"
// @Param aplication_id body int true "ApplicationID"
// @Param module_id body int true "ModuleID"
// @Param activity_id body int true "ActivityID"
// @Param user_id body string int "UserID"
// @Param description body string false "Description"
// @Success 200 {object} entities.ErrorLog
// @Security ApiKeyAuth
// @Router /v1/error-logs [post]
func CreateErrorLog(c *fiber.Ctx) error {
	// Create new error log struct
	errorlog := &entities.ErrorLog{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&errorlog); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a Company model.
	validate := utils.NewValidator()

	// Validate errorlog fields.
	if err := validate.Struct(errorlog); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Create errorlog by given model.
	if err := queries.CreateErrorLog(errorlog); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": errorlog,
		"meta": fiber.Map{},
	})
}

// DeleteErrorLog func for deletes error log by given ID.
// @Description Delete error log by given ID.
// @Summary delete error log by given ID
// @Tags ErrorLog
// @Accept json
// @Produce json
// @Param id body int true "ErrorLog ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/error-logs/{id} [delete]
func DeleteErrorLog(c *fiber.Ctx) error {
	// Catch error log ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Checking, if error log with given ID is exists.
	foundedErrorLog, err := queries.GetErrorLog(uint(id))
	if err != nil {
		// Return status 404 and company not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "error log with this ID not found",
		})
	}

	if err := queries.DeleteErrorLog(foundedErrorLog.ID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
