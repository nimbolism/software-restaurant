package database

import (
	"fmt"
	"log"
	"os"

	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConfig holds the configuration for the database connection.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var db *gorm.DB

// NewDBConfig creates a new DBConfig instance from environment variables.
func NewPQDBConfig() (*DBConfig, error) {
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
func OpenPQConnection(cfg *DBConfig) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Adjust logging level as needed
	})
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Ping the database
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to the PostgreSQL database: %v", err)
	}

	log.Println("Connected to the PostgreSQL database successfully!")

	// Run migrations
	if err := RunMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	return nil
}

// ClosePostgreSQLConnection closes the PostgreSQL database connection.
func ClosePQConnection() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return fmt.Errorf("failed to get underlying sql.DB: %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("failed to close PostgreSQL connection: %v", err)
		}
		log.Println("Closed PostgreSQL connection successfully!")
	}
	return nil
}

// RunMigrations runs migrations for the database.
func RunMigrations() error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Card{}, &models.Category{}, &models.Meal{}, &models.Food{}, &models.SideDish{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		return err
	}

	return nil
}

// GetDB returns the instance of *gorm.DB
func GetPQDB() *gorm.DB {
	return db
}
