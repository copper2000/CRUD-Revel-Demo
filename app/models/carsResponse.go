package models

type CarsResponse struct {
	Total int   `json:"total"`
	Car   []Car `json:"car"`
}
