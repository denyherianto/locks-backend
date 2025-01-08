package controllers

import (
	"strconv"
	"time"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/gofiber/fiber/v2"
)

// GetCompanies func gets all exists companies.
// @Description Get all exists companies.
// @Summary get all exists companies
// @Tags Companies
// @Accept json
// @Produce json
// @Success 200 {array} entities.Company
// @Router /v1/companies [get]
func GetCompanies(c *fiber.Ctx) error {
	// Get all companies.
	companies, err := queries.GetCompanies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": companies,
		"meta": fiber.Map{
			"total": len(companies),
		},
	})
}

// GetCompany func gets company by given ID or 404 error.
// @Description Get company by given ID.
// @Summary get company by given ID
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} entities.Company
// @Router /v1/companies/{id} [get]
func GetCompany(c *fiber.Ctx) error {
	// Catch company ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Get company by ID.
	company, err := queries.GetCompany(uint(id))
	if err != nil {
		// Return, if company not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "company with the given ID is not found",
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": company,
		"meta": fiber.Map{},
	})
}

// CreateCompany func for creates a new company.
// @Description Create a new company.
// @Summary create a new company
// @Tags Companies
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Success 200 {object} entities.Company
// @Security ApiKeyAuth
// @Router /v1/companies [post]
func CreateCompany(c *fiber.Ctx) error {
	// Create new Company struct
	company := &entities.Company{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(&company); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Create a new validator for a Company model.
	validate := utils.NewValidator()

	// Validate company fields.
	if err := validate.Struct(company); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Create company by given model.
	if err := queries.CreateCompany(company); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": company,
		"meta": fiber.Map{},
	})
}

// UpdateCompany func for updates company by given ID.
// @Description Update company.
// @Summary update company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id body string true "Company ID"
// @Param name body string true "Name"
// @Success 202 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/companies/{id} [put]
func UpdateCompany(c *fiber.Ctx) error {
	// Catch company ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Create a new user auth struct.
	company := &entities.Company{}

	// Checking received data from JSON body.
	if err := c.BodyParser(company); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Checking, if company with given ID is exists.
	foundedCompany, err := queries.GetCompany(uint(id))
	if err != nil {
		// Return status 404 and company not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "company with this ID not found",
		})
	}

	// Set initialized data for company:
	foundedCompany.Name = company.Name
	foundedCompany.UpdatedAt = time.Now()

	// Create a new validator for a Company model.
	validate := utils.NewValidator()
	// Validate company fields.
	if err := validate.Struct(foundedCompany); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   fiber.StatusBadRequest,
			"message": utils.ValidatorErrors(err),
		})
	}

	// Update company by given ID.
	if err := queries.UpdateCompany(uint(id), foundedCompany); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": foundedCompany,
		"meta": fiber.Map{},
	})
}

// DeleteCompany func for deletes company by given ID.
// @Description Delete company by given ID.
// @Summary delete company by given ID
// @Tags Companies
// @Accept json
// @Produce json
// @Param id body string true "Company ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/companies/{id} [delete]
func DeleteCompany(c *fiber.Ctx) error {
	// Catch company ID from URL.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Checking, if company with given ID is exists.
	foundedCompany, err := queries.GetCompany(uint(id))
	if err != nil {
		// Return status 404 and company not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "company with this ID not found",
		})
	}

	if err := queries.DeleteCompany(foundedCompany.ID); err != nil {
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
