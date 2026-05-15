package models

import "time"

// Transaction represents a sale record in the database
type Transaction struct {
	ID           int       `json:"id"`
	ProductID    int       `json:"product_id"`
	Quantity     int       `json:"quantity"`
	TotalPrice   int       `json:"total_price"`
	Status       string    `json:"status"`
	CustomerName string    `json:"customer_name"`
	CreatedAt    time.Time `json:"created_at"`
}
