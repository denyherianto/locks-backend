package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/app/utils"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetErrorLogs method for getting all error logs.
func GetErrorLogs(search string, pagination *utils.Pagination) ([]entities.ErrorLog, error) {
	// Define errorlogs variable.
	var errorlogs []entities.ErrorLog
	// Send query to database.
	result := database.DBManager.Scopes(utils.Paginate(&errorlogs, pagination, database.DBManager)).Where("description LIKE ?", "%"+search+"%").Find(&errorlogs)
	if result.Error != nil {
		// Return empty object and error.
		return nil, result.Error
	}
	// Return query result.
	return errorlogs, nil
}

// GetErrorLog method for getting one error log by given ID.
func GetErrorLog(id uint) (entities.ErrorLog, error) {
	// Define error log variable.
	var errorlog entities.ErrorLog
	// Send query to database.
	result := database.DBManager.First(&errorlog, id)
	if result.Error != nil {
		// Return empty object and error.
		return errorlog, result.Error
	}
	// Return query result.
	return errorlog, nil
}

// CreateErrorLog method for creating error log by given ErrorLog object.
func CreateErrorLog(errorlog *entities.ErrorLog) error {
	// Send query to database.
	result := database.DBManager.Create(errorlog)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// DeleteErrorLog method for delete error log by given ID.
func DeleteErrorLog(id uint) error {
	errorlog := entities.ErrorLog{}

	// Send query to database.
	result := database.DBManager.Delete(&errorlog, id)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}
