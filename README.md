# 📦 Go Inventory API

A high-performance, persistent RESTful API built with **Golang** to manage a warehouse inventory system. This project demonstrates core backend architecture principles, clean data modeling, and robust HTTP request handling without relying on heavy external frameworks.

[![Enterprise Go CI/CD](https://github.com/zonafirmann/go-inventory-api/actions/workflows/go-ci.yml/badge.svg)](https://github.com/zonafirmann/go-inventory-api/actions/workflows/go-ci.yml)
![Go Version](https://img.shields.io/github/go-mod/go-version/zonafirmann/go-inventory-api?style=flat&logo=go)
![License](https://img.shields.io/github/license/zonafirmann/go-inventory-api?style=flat)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)

A high-performance, persistent RESTful API... (lanjutkan teks aslinya)

## ⚡ Core Features
* **Full CRUD Operations:** Seamlessly Create, Read, Update, and Delete inventory items.
* **Persistent Storage:** Utilizes a custom JSON I/O engine to ensure data durability across server restarts.
* **Modular Architecture:** Clean separation of concerns with isolated struct models and routing logic.
* **Continuous Integration:** Automated build and testing pipeline powered by GitHub Actions.

## 🛠️ API Endpoints

| Method | Endpoint | Description | Payload (JSON) |
| :--- | :--- | :--- | :--- |
| `GET` | `/products` | Fetch all products in the warehouse | - |
| `POST` | `/products` | Add a new product | `{"name": "string", "stock": int, "price": int}` |
| `PUT` | `/products?id={id}` | Update stock and price of a product | `{"stock": int, "price": int}` |
| `DELETE`| `/products?id={id}` | Remove a product from the warehouse | - |

## 🛠️ Tech Stack Upgrade (2026 Standard)
* **Language:** Go 1.22+
* **Database:** PostgreSQL 18 (Relational Database)
* **Driver:** `pgx/v5` (High-performance PostgreSQL driver for Go)
* **CI/CD:** GitHub Actions with Static Analysis

## 🏗️ System Architecture
The system has been migrated from a flat JSON file to a robust **Relational Database Management System (RDBMS)**. This ensures:
1. **Data Integrity:** Using Primary Keys and strict data types.
2. **Concurrency:** Handling multiple requests without file locking issues.
3. **Security:** Implementation of **Prepared Statements** to prevent SQL Injection attacks.

## 🚀 How to Run Locally

1. Clone the repository and navigate to the project directory.
2. Run the server:
   ```bash
   go run main.go
