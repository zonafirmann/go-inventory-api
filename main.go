package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Connection string: Replace 'PASSWORD_KAMU' with your actual PostgreSQL password
	connString := "postgres://postgres:PASSWORD_KAMU@localhost:5432/inventory_db"

	// Establish connection to PostgreSQL 18
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Ensure connection is closed when the function finishes
	defer conn.Close(context.Background())

	fmt.Println("------------------------------------------")
	fmt.Println("✅ STATUS: CONNECTED TO POSTGRESQL 18")
	fmt.Println("------------------------------------------")

	// --- PART 1: INSERT DATA ---

	// Define new product data
	newName := "Gaming Mouse RGB"
	newStock := 25
	newPrice := 450000

	// SQL query using placeholders ($1, $2, $3) to prevent SQL Injection
	// This is a security best practice in 2026
	sqlStatement := `INSERT INTO products (name, stock, price) VALUES ($1, $2, $3)`

	// Execute the command
	_, err = conn.Exec(context.Background(), sqlStatement, newName, newStock, newPrice)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}

	fmt.Println("Successfully added new product to the database!")

	// Execute query to fetch all products from the table
	rows, err := conn.Query(context.Background(), "SELECT id, name, stock, price FROM products")
	if err != nil {
		log.Fatalf("Query execution failed: %v", err)
	}
	defer rows.Close()

	fmt.Println("📦 CURRENT INVENTORY LIST:")

	// Iterate through the result set
	for rows.Next() {
		var id int
		var name string
		var stock int
		var price int

		// Scan row data into local variables
		err := rows.Scan(&id, &name, &stock, &price)
		if err != nil {
			log.Fatalf("Row scan failed: %v", err)
		}

		// Print formatted output to the terminal
		fmt.Printf("ID: %d | Name: %-20s | Stock: %d | Price: Rp%d\n", id, name, stock, price)
	}

	// Check for any errors encountered during iteration
	if rows.Err() != nil {
		log.Fatalf("Error occurred during row iteration: %v", rows.Err())
	}
}
