package db

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "ITM-2020"
	dbname   = "pqgotest"
)

var DB *sql.DB

func InitDB() {
	
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname) // Example connection string
	var err error
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		panic(fmt.Errorf("could not connect to database: %w", err)) // Include error details
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        author_id TEXT NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL
    )
    `
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic(fmt.Errorf("could not create events table: %w", err)) // Include error details
	}

	
}
