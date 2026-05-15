package models

// Product represents a single item in our warehouse inventory.
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Price int    `json:"price"`
}
