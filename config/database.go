package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// ConnectDB establishes a connection to PostgreSQL 18
func ConnectDB() *pgx.Conn {
	// Connection string: Replace with your actual password
	connString := "postgres://postgres:yona20042006stsd*@localhost:5432/inventory_db"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Database connection established successfully")
	return conn
}
