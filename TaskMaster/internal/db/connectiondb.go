package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var database *sql.DB

// InitDB initializes the database connection
func InitDB() {
	var err error
	connStr := "user=taskmaster password=taskmaster dbname=taskmaster sslmode=disable"
	database, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Teste a conexão
	if err := database.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	)`
	if _, err = database.Exec(createUserTableQuery); err != nil {
		log.Fatal("Failed to create users table:", err)
	}

	createTaskTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		title VARCHAR(100) NOT NULL,
		status VARCHAR(50) NOT NULL
	)`
	if _, err = database.Exec(createTaskTableQuery); err != nil {
		log.Fatal("Failed to create tasks table:", err)
	}
}

func GetDB() *sql.DB {
	return database
}
