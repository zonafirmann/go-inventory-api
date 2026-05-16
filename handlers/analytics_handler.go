package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/zonafirmann/go-inventory-api/repository"
)

// GetAnalyticsHandler processes the inventory smart report request
func GetAnalyticsHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Restrict access to GET method only
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Generate the smart report from the repository layer
		report, err := repository.GetSmartAnalysis(db)
		if err != nil {
			http.Error(w, "Failed to generate report: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Set response headers and encode data to JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(report)
	}
}
