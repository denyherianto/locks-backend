package logger

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/queries"
)

func LogError(errorlog *entities.ErrorLog) error {
	// Send query to database.
	err := queries.CreateErrorLog(errorlog)
	if err != nil {
		// Return only error.
		return err
	}
	// This query returns nothing.
	return nil
}
