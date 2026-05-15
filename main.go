package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zonafirmann/go-inventory-api/config"
	"github.com/zonafirmann/go-inventory-api/repository"
)

func main() {
	// 1. Initialize Database Connection
	db := config.ConnectDB()
	defer db.Close(context.Background())

	// 2. Fetch all products using the repository layer
	products, err := repository.GetAllProducts(db)
	if err != nil {
		log.Fatalf("Failed to fetch products: %v", err)
	}

	// 3. Display Results
	fmt.Println("\n📦 GLOBAL INVENTORY REPORT:")
	for _, p := range products {
		fmt.Printf("[%d] %-20s | Stock: %-5d | Price: Rp%d\n", p.ID, p.Name, p.Stock, p.Price)
	}
}
