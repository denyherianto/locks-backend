package requests

import "github.com/denyherianto/go-fiber-boilerplate/app/models/entities"

type ErrorLogMeta struct {
	Limit      int `json:"limit"`
	Page       int `json:"page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ErrorLogResponse struct {
	Data []entities.ErrorLog `json:"data"` // Data berisi array objek Item
	Meta ErrorLogMeta        `json:"meta"`
}
