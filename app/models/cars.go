package models

type Car struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Engine string `json:"engine"`
}
