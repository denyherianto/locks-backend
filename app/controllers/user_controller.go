package controllers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
)

// GetUsersByIDs func gets all exists companies.
// @Description Get all exists companies.
// @Summary get all exists companies
// @Tags Users
// @Accept json
// @Produce json
// @Param ids path string true "User IDs"
// @Success 200 {array} entities.User
// @Router /v1/users [get]
func GetUsersByIDs(c *fiber.Ctx) error {
	q := c.Queries()

	ids := make([]uint, len(q["ids"]))
	for i, stringId := range strings.Split(q["ids"], ",") {
		parsed, _ := strconv.Atoi(stringId)
		ids[i] = uint(parsed)
	}

	// Get company by ID.
	users, err := queries.GetUserByIDs(ids)
	if err != nil {
		// Return, if company not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   fiber.StatusNotFound,
			"message": "users with the given ID is not found",
		})
	}

	usersMap := make(map[string]entities.User)

	for _, user := range *users {
		usersMap[strconv.Itoa(int(user.ID))] = user
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"data": usersMap,
		"meta": fiber.Map{},
	})
}
