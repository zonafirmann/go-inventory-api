package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// InventoryReport represents the smart analysis result for a product
type InventoryReport struct {
	ProductID      int    `json:"product_id"`
	ProductName    string `json:"product_name"`
	CurrentStock   int    `json:"current_stock"`
	StatusAlert    string `json:"status_alert"` // CRITICAL, OVERSTOCK, or NORMAL
	Recommendation string `json:"recommendation"`
}

// GetSmartAnalysis analyzes database records without external AI APIs
func GetSmartAnalysis(conn *pgx.Conn) ([]InventoryReport, error) {
	ctx := context.Background()

	// Fetch all products to analyze their current state
	rows, err := conn.Query(ctx, "SELECT id, name, stock, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []InventoryReport

	// Iterate through each product to apply smart business rules
	for rows.Next() {
		var id, stock, price int
		var name string

		if err := rows.Scan(&id, &name, &stock, &price); err != nil {
			return nil, err
		}

		// Default healthy parameters
		status := "NORMAL"
		recommendation := "Stock level is healthy. No action required."

		// Rule 1: Check for critical low stock
		if stock <= 3 {
			status = "CRITICAL ALERT"
			recommendation = "Stock is dangerously low! Reorder immediately from supplier."
			// Rule 2: Check for overstock on high-value items
		} else if stock > 50 && price > 1000000 {
			status = "OVERSTOCK WARNING"
			recommendation = "Capital tied up in expensive inventory. Consider a promotional discount."
		}

		// Append the analysis result to the report slice
		reports = append(reports, InventoryReport{
			ProductID:      id,
			ProductName:    name,
			CurrentStock:   stock,
			StatusAlert:    status,
			Recommendation: recommendation,
		})
	}

	return reports, nil
}
