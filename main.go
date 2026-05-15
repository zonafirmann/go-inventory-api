package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/zonafirmann/go-inventory-api/config"
	"github.com/zonafirmann/go-inventory-api/handlers"
)

func main() {
	// 1. Initialize Database Connection
	db := config.ConnectDB()
	defer db.Close(context.Background())

	// 2. Define API Routes (HARUS SEBELUM ListenAndServe)
	http.HandleFunc("/products", handlers.GetProductsHandler(db))
	http.HandleFunc("/checkout", handlers.CheckoutHandler(db)) // <--- PINDAH KE SINI

	// 3. Start the Web Server
	port := ":8080"
	fmt.Printf("🚀 Inventory API Server is running on http://localhost%s\n", port)

	// Baris ini harus menjadi baris TERAKHIR yang mengeksekusi rute
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
