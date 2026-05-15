package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/zonafirmann/go-inventory-api/repository"
)

type CheckoutRequest struct {
	ProductID    int    `json:"product_id"`
	Quantity     int    `json:"quantity"`
	CustomerName string `json:"customer_name"`
}

func CheckoutHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req CheckoutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Call the repository to process the transaction
		err := repository.CreateTransaction(db, req.ProductID, req.Quantity, req.CustomerName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Transaction successful!"})
	}
}
