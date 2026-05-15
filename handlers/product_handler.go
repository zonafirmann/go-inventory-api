package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/zonafirmann/go-inventory-api/repository"
)

// GetProductsHandler handles the GET /products request
func GetProductsHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Fetch data from the repository layer
		products, err := repository.GetAllProducts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// 2. Set response header to JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// 3. Send the products data as JSON response
		json.NewEncoder(w).Encode(products)
	}
}
