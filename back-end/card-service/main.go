package main

import (
	"log"

	"github.com/nimbolism/software-restaurant/back-end/database"
)

func main() {
	// Initialize the PostgreSQL database connection
	cfg, err := database.NewDBConfig()
	if err != nil {
		log.Fatalf("Failed to create DBConfig: %v", err)
	}

	// Open a connection to the PostgreSQL database
	db, err := database.OpenPostgreSQLConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to open PostgreSQL connection: %v", err)
	}
	defer func() {
		if err := database.ClosePostgreSQLConnection(db); err != nil {
			log.Fatalf("Failed to close PostgreSQL connection: %v", err)
		}
	}()

	// Use db to perform a database operation
	// For example, query the users table
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatalf("Failed to query database: %v", err)
	}
	defer rows.Close()

	println("from the second service!")
}
