# 📦 Go Inventory API (PostgreSQL Edition)

A high-performance Warehouse Inventory system built with Golang and PostgreSQL 18. This project demonstrates professional backend architecture, secure data persistence, and RESTful API standards.

## 🚀 Features
* **Clean Architecture:** Separated layers for Config, Models, Repository, and Handlers.
* **Relational Database:** Powered by **PostgreSQL 18** for scalable data management.
* **Security:** Implemented **Prepared Statements** to prevent SQL Injection.
* **RESTful Endpoints:** Standardized JSON responses for frontend integration.

## 🏗️ System Architecture
The application follows the **Separation of Concerns** principle:
- **`config/`**: Database connection management.
- **`models/`**: Data structures and JSON mapping.
- **`repository/`**: Direct SQL interactions with PostgreSQL.
- **`handlers/`**: HTTP request and response logic.

## 🛠️ API Documentation
| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/products` | Fetch all products from the database |
| `GET` | `/products/analytics` | Fetch automated smart inventory analysis and stock alerts |

## ⚙️ Prerequisites
- Go 1.22+
- PostgreSQL 18
- `pgx/v5` Driver

## 🚦 Getting Started
1. Clone the repository.
2. Configure your PostgreSQL connection in `config/database.go`.
3. Run `go run main.go`.
4. Access the API at `http://localhost:8080/products`.

## 💸 Transaction Module (Mini Project Phase 1)
The API now supports secure transaction processing with the following features:
- **ACID Transactions:** Using `Begin`, `Commit`, and `Rollback` to ensure data consistency.
- **Stock Validation:** Prevents sales if stock is insufficient.
- **Atomic Operations:** Automatically updates product stock upon successful transaction.

## 💡 Business Value & Smart Features
This API doesn't just store data; it provides actionable business intelligence:
- **Smart Analytics Engine:** Rule-based algorithm to detect critical stock levels and overstock inefficiencies without relying on paid third-party APIs.
- **Financial Integrity:** ACID-compliant checkout system ensuring exact stock deductions and preventing race conditions during high-traffic sales.

### New API Endpoint
| Method | Endpoint | Description | Payload (JSON) |
| :--- | :--- | :--- | :--- |
| `POST` | `/checkout` | Process a product purchase | `{"product_id": 1, "quantity": 2, "customer_name": "Zona"}` |
---
**Status:** Core Backend Foundation Completed. Moving towards Fintech (Midtrans) Integration.