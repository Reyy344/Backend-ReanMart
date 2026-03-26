package models

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Ratings     float64 `json:"ratings"`
	ImageURL    string  `json:"image_url"`
	IsAvailable bool    `json:"is_available"`
	TotalSold   float64 `json:"total_sold"`
}