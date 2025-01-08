package controllers

import (
	"strconv"
	"time"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
)

// GetKeys func gets all exists keys.
// @Description Get all exists keys.
// @Summary get all exists keys
// @Tags Keys
// @Accept json
// @Produce json
// @Success 200 {array} entities.Key
// @Router /v1/keys [get]
func GetKeys(c *fiber.Ctx) error {
	search := c.Query("search", "")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	// Create pagination object.
	pagination := &utils.Pagination{
		Page:  page,
		Limit: limit,
	}

	// Get all keys.
	keys, err := queries.GetKeys(search, pagination)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": keys,
		"meta": fiber.Map{
			"page":       pagination.Page,
			"total_page": pagination.TotalPages,
			"limit":      pagination.Limit,
			"total":      pagination.TotalRows,
		},
	})
}

// GetKey func gets key by given ID or 404 error.
// @Description Get key by given ID.
// @Summary get key by given ID
// @Tags Keys
// @Accept json
// @Produce json
// @Param id path string true "Key ID"
// @Success 200 {object} entities.Key
// @Router /v1/keys/{id} [get]
func GetKey(c *fiber.Ctx) error {
	// Catch key ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Get key by ID.
	key, err := queries.GetKey(uint(id))
	if err != nil {
		// Return, if key not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "key with the given ID is not found",
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": key,
		"meta": fiber.Map{},
	})
}

// CreateKey func for creates a new key.
// @Description Create a new key.
// @Summary create a new key
// @Tags Keys
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Success 200 {object} entities.Key
// @Security ApiKeyAuth
// @Router /v1/keys [post]
func CreateKey(c *fiber.Ctx) error {
	// Create new Key struct
	key := &entities.Key{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&key); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a Key model.
	validate := utils.NewValidator()

	// Validate key fields.
	if err := validate.Struct(key); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Create key by given model.
	if err := queries.CreateKey(key); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	if key.CompanyID != 0 {
		// Get company by ID.
		company, err := queries.GetCompany(key.CompanyID)
		if err != nil {
			// Return, if company not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   fiber.StatusNotFound,
				"message": "company with the given ID is not found",
			})
		}

		key.Company = &company
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": key,
		"meta": fiber.Map{},
	})
}

// UpdateKey func for updates key by given ID.
// @Description Update key.
// @Summary update key
// @Tags Keys
// @Accept json
// @Produce json
// @Param id body string true "Key ID"
// @Param name body string true "Name"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/keys/{id} [put]
func UpdateKey(c *fiber.Ctx) error {
	// Catch key ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Create a new user auth struct.
	key := &entities.Key{}

	// Checking received data from JSON body.
	if err := c.BodyParser(key); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Checking, if key with given ID is exists.
	foundedKey, err := queries.GetKey(uint(id))
	if err != nil {
		// Return status 404 and key not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "key with this ID not found",
		})
	}

	// Set initialized data for key:
	foundedKey.UpdatedAt = time.Now()

	if key.Name != "" {
		foundedKey.Name = key.Name
	}

	if key.Description != "" {
		foundedKey.Description = key.Description
	}

	if key.CompanyID != 0 {
		// Get company by ID.
		company, err := queries.GetCompany(key.CompanyID)
		if err != nil {
			// Return, if company not found.
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   fiber.StatusNotFound,
				"message": "company with the given ID is not found",
			})
		}

		foundedKey.CompanyID = key.CompanyID
		foundedKey.Company = &company
	}

	// Create a new validator for a Key model.
	validate := utils.NewValidator()
	// Validate key fields.
	if err := validate.Struct(foundedKey); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Update key by given ID.
	if err := queries.UpdateKey(uint(id), foundedKey); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": foundedKey,
		"meta": fiber.Map{},
	})
}

// DeleteKey func for deletes key by given ID.
// @Description Delete key by given ID.
// @Summary delete key by given ID
// @Tags Keys
// @Accept json
// @Produce json
// @Param id body string true "Key ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/keys/{id} [delete]
func DeleteKey(c *fiber.Ctx) error {
	// Catch key ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Checking, if key with given ID is exists.
	foundedKey, err := queries.GetKey(uint(id))
	if err != nil {
		// Return status 404 and key not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "key with this ID not found",
		})
	}

	if err := queries.DeleteKey(foundedKey.ID); err != nil {
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
