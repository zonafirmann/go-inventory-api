package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

// CreateTransaction handles the checkout logic with atomicity
func CreateTransaction(conn *pgx.Conn, productID int, qty int, customer string) error {
	ctx := context.Background()

	// 1. Start a database transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	// Ensure rollback if something goes wrong
	defer tx.Rollback(ctx)

	// 2. Check current stock and price
	var stock int
	var price int
	err = tx.QueryRow(ctx, "SELECT stock, price FROM products WHERE id = $1 FOR UPDATE", productID).Scan(&stock, &price)
	if err != nil {
		return errors.New("product not found")
	}

	// 3. Validation: Is stock enough?
	if stock < qty {
		return errors.New("insufficient stock")
	}

	// 4. Update product stock
	newStock := stock - qty
	_, err = tx.Exec(ctx, "UPDATE products SET stock = $1 WHERE id = $2", newStock, productID)
	if err != nil {
		return err
	}

	// 5. Create transaction record
	totalPrice := price * qty
	_, err = tx.Exec(ctx,
		"INSERT INTO transactions (product_id, quantity, total_price, customer_name, status) VALUES ($1, $2, $3, $4, $5)",
		productID, qty, totalPrice, customer, "success")
	if err != nil {
		return err
	}

	// 6. Commit the transaction to the database
	return tx.Commit(ctx)
}
