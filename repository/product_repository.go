package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/zonafirmann/go-inventory-api/models"
)

// GetAllProducts fetches all inventory records from PostgreSQL
func GetAllProducts(conn *pgx.Conn) ([]models.Product, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, name, stock, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Stock, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
