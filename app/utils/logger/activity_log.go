package logger

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"

	"github.com/gofiber/fiber/v2"
)

func LogActivity(c *fiber.Ctx, activity *entities.Activity) error {
	activity.IPAddress = c.IP()
	activity.UserAgent = string(c.Request().Header.UserAgent())

	// Send query to database.
	err := queries.CreateActivity(activity)
	if err != nil {
		// Return only error.
		return err
	}
	// This query returns nothing.
	return nil
}
