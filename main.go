package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// 1. String Koneksi (URL Format)
	// Format: postgres://username:password@localhost:5432/nama_database
	connString := "postgres://postgres:yona20042006stsd*@localhost:5432/inventory_db"

	// 2. Membuka koneksi ke PostgreSQL 18
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal koneksi ke database: %v\n", err)
		os.Exit(1)
	}

	// Pastikan koneksi ditutup saat program selesai
	defer conn.Close(context.Background())

	// 3. Tes Koneksi dengan Query sederhana
	var productName string
	var price int
	err = conn.QueryRow(context.Background(), "SELECT name, price FROM products WHERE id=1").Scan(&productName, &price)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query gagal: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("✅ STATUS: BERHASIL TERKONEKSI KE POSTGRES 18")
	fmt.Printf("📦 Data Produk Pertama: %s (Harga: Rp%d)\n", productName, price)
	fmt.Println("------------------------------------------")
}
