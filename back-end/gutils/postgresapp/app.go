package postgresapp

import (
	"log"

	"github.com/nimbolism/software-restaurant/back-end/database"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

var DB *gorm.DB

func New() *App {
	cfg, err := database.NewPQDBConfig()
	if err != nil {
		log.Fatalf("Failed to create DBConfig: %v", err)
	}

	err = database.OpenPQConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to open PostgreSQL connection: %v", err)
	}

	DB = database.GetPQDB()

	app := &App{
		DB: DB,
	}

	return app
}

func (a *App) Close() {
	if err := database.ClosePQConnection(); err != nil {
		log.Fatalf("Error closing PostgreSQL connection: %v", err)
	}
}
