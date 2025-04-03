# GO TUTORIAL
## By Raddames Tonui

### Summary
This tutorial provides a step-by-step guide to building an e-commerce API using Go and the Gin framework. It covers project setup, database connection, migrations, and CRUD API implementation.

## Project Structure
```
	go-ecommerce-api/
	│── main.go
	│── migrations/
	│   ├── 000001_init_schema.down.sql
	│   ├── 000001_init_schema.up.sql
	│── database/
	│   ├── database.go
	│── models/
	│   ├── user.go
	│   ├── product.go
	│   ├── order.go
	│── routes/
	│   ├── user_routes.go
	│   ├── product_routes.go
	│   ├── order_routes.go
	│── controllers/
	│   ├── user_controller.go
	│   ├── product_controller.go
	│   ├── order_controller.go
	│── .env
	│── go.mod
	│── go.sum
	│── README.md
	
```

## Steps to Build the API

### Step 1: Setting up Go Project
[Go Environment Setup](documentation/1_GoSetUp.md)

### Step 2: Setting Up Database Connection
[Database Connection Documentation](documentation/2_DabaseConnection.md)

### Step 3: Generate and Run Database Migrations
[Database Migrations Documentation](documentation/3_DatabaseMigrations.md)

### Step 4: Building API Endpoints – Implementing CRUD Routes
[API CRUD Documentation](documentation/4_GoGinCrudApi.md)

---
### Connect with Me
[![LinkedIn](https://img.shields.io/badge/LinkedIn-blue?style=for-the-badge)](https://www.linkedin.com/in/raddames-tonui-01a751277/)


URL Shortener – Learn routing, database storage, and hashing techniques.

Blog API – Practice CRUD operations, authentication, and relational databases.

Chat Application (WebSockets) – Explore real-time communication.

E-commerce API – Implement authentication, payments, and order management.

Weather API Fetcher – Work with external APIs, JSON parsing, and caching.

Expense Tracker – Store user expenses, categorize them, and generate reports.