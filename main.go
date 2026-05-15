package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/zonafirmann/go-inventory-api/models"
)

const fileName = "products.json"

// LoadProducts reads the JSON database and returns a slice of products.
func LoadProducts() []models.Product {
	var products []models.Product
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return products
	}
	json.Unmarshal(bytes, &products)
	return products
}

// SaveProducts writes the slice of products back to the JSON database.
func SaveProducts(products []models.Product) {
	bytes, _ := json.MarshalIndent(products, "", "  ")
	os.WriteFile(fileName, bytes, 0644)
}

// ---------------------------------------------------------
// CONTROLLER: Handles Full CRUD for Inventory System
// ---------------------------------------------------------
func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// READ: Return all products in the warehouse
		inventory := LoadProducts()
		json.NewEncoder(w).Encode(inventory)
		fmt.Println("[LOG] GET request processed: Displaying inventory")

	case http.MethodPost:
		// CREATE: Add a new product to the warehouse
		var newProduct models.Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		inventory := LoadProducts()
		newProduct.ID = len(inventory) + 1

		inventory = append(inventory, newProduct)
		SaveProducts(inventory)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newProduct)
		fmt.Printf("[LOG] POST request processed: Added %s to inventory\n", newProduct.Name)

	case http.MethodPut:
		// UPDATE: Modify stock or price of an existing product via JSON payload
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || idStr == "" {
			http.Error(w, "Invalid or missing product ID", http.StatusBadRequest)
			return
		}

		var updatedData models.Product
		err = json.NewDecoder(r.Body).Decode(&updatedData)
		if err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		inventory := LoadProducts()
		productFound := false

		for i, product := range inventory {
			if product.ID == id {
				// Update only the Stock and Price fields
				inventory[i].Stock = updatedData.Stock
				inventory[i].Price = updatedData.Price
				productFound = true
				break
			}
		}

		if !productFound {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		SaveProducts(inventory)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Product successfully updated"}`))
		fmt.Printf("[LOG] PUT request processed: Updated product %d\n", id)

	case http.MethodDelete:
		// DELETE: Remove a product from the warehouse
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || idStr == "" {
			http.Error(w, "Invalid or missing product ID", http.StatusBadRequest)
			return
		}

		inventory := LoadProducts()
		productFound := false

		for i, product := range inventory {
			if product.ID == id {
				// Remove the item from the slice
				inventory = append(inventory[:i], inventory[i+1:]...)
				productFound = true
				break
			}
		}

		if !productFound {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		SaveProducts(inventory)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Product successfully deleted"}`))
		fmt.Printf("[LOG] DELETE request processed: Removed product %d\n", id)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	fmt.Println("=== WAREHOUSE INVENTORY API STARTED ===")
	http.HandleFunc("/products", productsHandler)
	port := ":8080"
	fmt.Printf("[INFO] Server is listening on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("[FATAL ERROR]", err)
	}
}
