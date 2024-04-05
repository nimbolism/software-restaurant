package main

import (
	"log"

	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/user-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http"
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

	// Start gRPC server
	go grpc.StartServer()

	// Start HTTP server
	go http.StartServer(db)

	// Keep the main function running
	select {}
}
