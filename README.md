# 📦 Go Inventory API

A high-performance, persistent RESTful API built with **Golang** to manage a warehouse inventory system. This project demonstrates core backend architecture principles, clean data modeling, and robust HTTP request handling without relying on heavy external frameworks.

# 📦 Go Inventory API

[![Enterprise Go CI/CD](https://github.com/zonafirmann/go-inventory-api/actions/workflows/go-ci.yml/badge.svg)](https://github.com/zonafirmann/go-inventory-api/actions/workflows/go-ci.yml)
![Go Version](https://img.shields.io/github/go-mod/go-version/zonafirmann/go-inventory-api?style=flat&logo=go)
![License](https://img.shields.io/github/license/zonafirmann/go-inventory-api?style=flat)

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

## 🚀 How to Run Locally

1. Clone the repository and navigate to the project directory.
2. Run the server:
   ```bash
   go run main.go
