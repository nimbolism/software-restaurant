package database

import (
	"fmt"
	"log"
	"os"
	"time"

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
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tehran",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // Adjust logging level as needed
		})
		if err != nil {
			log.Printf("Failed to open PostgreSQL connection, retrying...: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Println("Failed to get underlying sql.DB, retrying...")
			time.Sleep(time.Second * 5)
			continue
		}
		if err := sqlDB.Ping(); err != nil {
			log.Println("Failed to connect to the PostgreSQL database, retrying...")
			time.Sleep(time.Second * 5)
			continue
		}

		log.Println("Connected to the PostgreSQL database successfully!")

		if err := RunMigrations(); err != nil {
			return fmt.Errorf("failed to run migrations: %v", err)
		}

		return nil
	}

	return fmt.Errorf("failed to connect to the PostgreSQL database after 10 retries")
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
	models := []interface{}{
		&models.User{},
		&models.Card{},
		&models.Category{},
		&models.Meal{},
		&models.Food{},
		&models.SideDish{},
		&models.Order{},
		&models.OrderFail{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}

// // GetPQDB returns the instance of *gorm.DB
func GetPQDB() *gorm.DB {
	return db
}
