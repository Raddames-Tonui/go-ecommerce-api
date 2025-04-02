# Step 3: Generate and Run Database Migrations

## There are two ways to perform database migrations:

### a) Using GORM (ORM-based migrations)
GORM provides an **AutoMigrate** feature that automatically creates and updates database tables based on Go structs.

### b) Using SQL Migrate (Raw SQL migrations)
SQL-based migrations give you full control over schema changes and allow you to define exact SQL queries.

---

## 3.1: Using GORM AutoMigrate

### 3.1.1: Install GORM
Ensure you have **GORM** and the PostgreSQL driver installed:

```sh
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### 3.1.2: Define Models with Relationships

GORM automatically handles relationships and timestamps when `gorm.Model` is embedded.

#### `models/user.go`
```go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username  string  `json:"username" gorm:"not null;unique"`
    Email     string  `json:"email" gorm:"not null;unique"`
    Password  string  `json:"password" gorm:"not null"`
    FirstName string  `json:"first_name"`
    LastName  string  `json:"last_name"`
    Orders    []Order `json:"orders" gorm:"foreignKey:UserID"`
}
```

#### `models/product.go`
```go
package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name        string  `json:"name" gorm:"not null"`
    Description string  `json:"description"`
    Price       float64 `json:"price" gorm:"not null"`
    Stock       int     `json:"stock" gorm:"not null"`
    Orders      []Order `gorm:"foreignKey:ProductID"`
}
```

#### `models/order.go`
```go
package models

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    UserID     uint    `json:"user_id" gorm:"not null"`
    User       User    `gorm:"foreignKey:UserID"`
    ProductID  uint    `json:"product_id" gorm:"not null"`
    Product    Product `gorm:"foreignKey:ProductID"`
    Quantity   int     `json:"quantity" gorm:"not null"`
    TotalPrice float64 `json:"total_price" gorm:"not null"`
}
```

### 3.1.3: Automate Migrations in `database/migrations.go`
```go
    package database

    import (
        "fmt"
        "log"
        "os"

        "go-ecommerce-api/models"

        "github.com/joho/godotenv"
        "gorm.io/driver/postgres"
        "gorm.io/gorm"
    )

    func MigrateDb () *gorm.DB {
        // Load environment variables from .env file
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }

        // Build database connection string
        dsn := fmt.Sprintf(
            "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
            os.Getenv("DB_HOST"),
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASSWORD"),
            os.Getenv("DB_NAME"),
            os.Getenv("DB_PORT"), 
        )

        // Connect to the database
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatal("Failed to connect to database: ", err)
        }
        
        fmt.Println("Database connected successfully!")
        // Migrate the schema
        err = db.AutoMigrate(
            &models.User{},
            &models.Product{},
            &models.Order{},
        )
        if err != nil {
            log.Fatal("Failed to migrate database: ", err)
        }
        fmt.Println("Database connected and migrated successfully!")
        return db
    }
```

### 3.1.4: Run the Migrations

When you start the application, GORM will automatically create the tables if they don't exist.

```sh
go run main.go
```

---

## 3.2: Using SQL Migrate (Raw SQL Migrations)

### 3.2.1: Install golang-migrate

```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Ensure it's installed correctly:
```sh
migrate -version
```

### 3.2.2: Create a Migrations Directory

```sh
mkdir migrations
```

### 3.2.3: Generate a New Migration File

```sh
migrate create -ext sql -dir migrations -seq init_schema
```

This creates two files:
```
migrations/
  000001_init_schema.up.sql
  000001_init_schema.down.sql
```

### 3.2.4: Define the Migration Queries

#### `migrations/000001_init_schema.up.sql`
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
```

#### `migrations/000001_init_schema.down.sql`
```sql
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;
```

### 3.2.5: Run Migrations

```sh
migrate -path migrations -database "postgres://admin:admin@localhost:5432/ecommerce_db?sslmode=disable" up
```

To rollback:
```sh
migrate -path migrations -database "postgres://admin:admin@localhost:5432/ecommerce_db?sslmode=disable" down
```

### 3.2.6: Verify Migration

```sql
SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';
```

---

## 3.3: Conclusion
- **GORM AutoMigrate** is simpler but offers less control.
- **SQL Migrations** give full control over schema changes.

📖 **For more documentation on GORM vs. raw SQL, refer to:**
[SQL_vs_GORM.md](documentations/SQL_vs_GORM.md)

