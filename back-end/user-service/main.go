package main

import (
	"log"

	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/user-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http"
)

func main() {
	// Initialize the PostgreSQL database connection
	cfg, err := database.NewPQDBConfig()
	if err != nil {
		log.Fatalf("Failed to create DBConfig: %v", err)
	}

	// Open a connection to the PostgreSQL database
	err = database.OpenPQConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to open PostgreSQL connection: %v", err)
	}
	defer func() {
		if err := database.ClosePQConnection(); err != nil {
			log.Fatalf("Error closing PostgreSQL connection: %v", err)
		}
	}()

	// Start gRPC server
	go grpc.StartServer()

	// Start HTTP server
	go http.StartServer()

	// Keep the main function running
	select {}
}
