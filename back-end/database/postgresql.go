package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// DBConfig holds the configuration for the database connection.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewDBConfig creates a new DBConfig instance from environment variables.
func NewDBConfig() (*DBConfig, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || password == "" || dbName == "" {
		return nil, fmt.Errorf("one or more PostgreSQL environment variables are not set")
	}

	return &DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}, nil
}

// OpenPostgreSQLConnection opens a connection to the PostgreSQL database.
func OpenPostgreSQLConnection(cfg *DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to the PostgreSQL database: %v", err)
	}

	log.Println("Connected to the PostgreSQL database successfully!")
	return db, nil
}

// ClosePostgreSQLConnection closes the PostgreSQL database connection.
func ClosePostgreSQLConnection(db *sql.DB) error {
	if db != nil {
		if err := db.Close(); err != nil {
			return fmt.Errorf("failed to close PostgreSQL connection: %v", err)
		}
		log.Println("Closed PostgreSQL connection successfully!")
	}
	return nil
}
