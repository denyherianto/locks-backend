package queries

import (
	"github.com/denyherianto/go-fiber-boilerplate/app/models/entities"
	"github.com/denyherianto/go-fiber-boilerplate/configs/database"
)

// GetCompanies method for getting all companies.
func GetCompanies() ([]entities.Company, error) {
	// Define companies variable.
	var companies []entities.Company
	// Send query to database.
	result := database.DBManager.Find(&companies)
	if result.Error != nil {
		// Return empty object and error.
		return nil, result.Error
	}
	// Return query result.
	return companies, nil
}

// GetCompany method for getting one company by given ID.
func GetCompany(id uint) (entities.Company, error) {
	// Define company variable.
	var company entities.Company
	// Send query to database.
	result := database.DBManager.First(&company, id)
	if result.Error != nil {
		// Return empty object and error.
		return company, result.Error
	}
	// Return query result.
	return company, nil
}

// CreateCompany method for creating company by given Company object.
func CreateCompany(company *entities.Company) error {
	// Send query to database.
	result := database.DBManager.Create(company)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// UpdateCompany method for updating company by given Company object.
func UpdateCompany(id uint, company entities.Company) error {
	// Send query to database.
	result := database.DBManager.Where("id = ?", id).Updates(&company)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}

// DeleteCompany method for delete company by given ID.
func DeleteCompany(id uint) error {
	company := entities.Company{}

	// Send query to database.
	result := database.DBManager.Delete(&company, id)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}
	// This query returns nothing.
	return nil
}
