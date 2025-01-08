package controllers

import (
	"strconv"
	"time"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
)

// GetKeyCopies func gets all exists keyCopies.
// @Description Get all exists keyCopies.
// @Summary get all exists keyCopies
// @Tags KeyCopies
// @Accept json
// @Produce json
// @Success 200 {array} entities.KeyCopy
// @Router /v1/keyCopies [get]
func GetKeyCopies(c *fiber.Ctx) error {
	search := c.Query("search", "")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	// Create pagination object.
	pagination := &utils.Pagination{
		Page:  page,
		Limit: limit,
	}

	// Get all keyCopies.
	keyCopies, err := queries.GetKeyCopies(search, pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": keyCopies,
		"meta": fiber.Map{
			"page":       pagination.Page,
			"total_page": pagination.TotalPages,
			"limit":      pagination.Limit,
			"total":      pagination.TotalRows,
		},
	})
}

// GetKeyCopy func gets keyCopy by given ID or 404 error.
// @Description Get keyCopy by given ID.
// @Summary get keyCopy by given ID
// @Tags KeyCopies
// @Accept json
// @Produce json
// @Param id path string true "KeyCopy ID"
// @Success 200 {object} entities.KeyCopy
// @Router /v1/keyCopies/{id} [get]
func GetKeyCopy(c *fiber.Ctx) error {
	// Catch keyCopy ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Get keyCopy by ID.
	keyCopy, err := queries.GetKeyCopy(uint(id))
	if err != nil {
		// Return, if keyCopy not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "keyCopy with the given ID is not found",
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": keyCopy,
		"meta": fiber.Map{},
	})
}

// CreateKeyCopy func for creates a new keyCopy.
// @Description Create a new keyCopy.
// @Summary create a new keyCopy
// @Tags KeyCopies
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Success 200 {object} entities.KeyCopy
// @Security ApiKeyCopyAuth
// @Router /v1/keyCopies [post]
func CreateKeyCopy(c *fiber.Ctx) error {
	// Create new KeyCopy struct
	keyCopy := &entities.KeyCopy{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&keyCopy); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a KeyCopy model.
	validate := utils.NewValidator()

	// Validate keyCopy fields.
	if err := validate.Struct(keyCopy); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Create keyCopy by given model.
	if err := queries.CreateKeyCopy(keyCopy); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	
	if keyCopy.KeyID != 0 {
		// Get key by ID.
		key, err := queries.GetKey(keyCopy.KeyID)
		if err != nil {
			// Return, if key not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   fiber.StatusNotFound,
				"message": "key with the given ID is not found",
			})
		}

		keyCopy.Key = &key
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": keyCopy,
		"meta": fiber.Map{},
	})
}

// UpdateKeyCopy func for updates keyCopy by given ID.
// @Description Update keyCopy.
// @Summary update keyCopy
// @Tags KeyCopies
// @Accept json
// @Produce json
// @Param id body string true "KeyCopy ID"
// @Param name body string true "Name"
// @Success 202 {string} status "ok"
// @Security ApiKeyCopyAuth
// @Router /v1/keyCopies/{id} [put]
func UpdateKeyCopy(c *fiber.Ctx) error {
	// Catch keyCopy ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Create a new user auth struct.
	keyCopy := &entities.KeyCopy{}

	// Checking received data from JSON body.
	if err := c.BodyParser(keyCopy); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Checking, if keyCopy with given ID is exists.
	foundedKeyCopy, err := queries.GetKeyCopy(uint(id))
	if err != nil {
		// Return status 404 and keyCopy not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "keyCopy with this ID not found",
		})
	}

	// Set initialized data for keyCopy:
	foundedKeyCopy.UpdatedAt = time.Now()

	if foundedKeyCopy.Identifier != "" {
		foundedKeyCopy.Identifier = keyCopy.Identifier
	}

	if foundedKeyCopy.Status != "" {
		foundedKeyCopy.Status = keyCopy.Status
	}

	if keyCopy.KeyID != 0 {
		// Get key by ID.
		key, err := queries.GetKey(keyCopy.KeyID)
		if err != nil {
			// Return, if key not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   fiber.StatusNotFound,
				"message": "key with the given ID is not found",
			})
		}

		foundedKeyCopy.KeyID = keyCopy.KeyID
		foundedKeyCopy.Key = &key
	}

	// Create a new validator for a KeyCopy model.
	validate := utils.NewValidator()
	// Validate keyCopy fields.
	if err := validate.Struct(foundedKeyCopy); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Update keyCopy by given ID.
	if err := queries.UpdateKeyCopy(uint(id), foundedKeyCopy); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": foundedKeyCopy,
		"meta": fiber.Map{},
	})
}

// DeleteKeyCopy func for deletes keyCopy by given ID.
// @Description Delete keyCopy by given ID.
// @Summary delete keyCopy by given ID
// @Tags KeyCopies
// @Accept json
// @Produce json
// @Param id body string true "KeyCopy ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyCopyAuth
// @Router /v1/keyCopies/{id} [delete]
func DeleteKeyCopy(c *fiber.Ctx) error {
	// Catch keyCopy ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Checking, if keyCopy with given ID is exists.
	foundedKeyCopy, err := queries.GetKeyCopy(uint(id))
	if err != nil {
		// Return status 404 and keyCopy not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "keyCopy with this ID not found",
		})
	}

	if err := queries.DeleteKeyCopy(foundedKeyCopy.ID); err != nil {
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
