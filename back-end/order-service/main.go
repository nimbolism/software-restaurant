package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/order-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/order-service/http"
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

	// Initialize the gRPC client connection
	if err := grpc.InitializeUserGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	if err := grpc.InitializeCardGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	if err := grpc.InitializeFoodGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	// Start gRPC server
	go grpc.StartServer()

	// Start HTTP server
	go http.StartServer()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		// Close the gRPC client connection when the server is shutting down
		grpc.CloseUserGRPCClient()
		grpc.CloseCardGRPCClient()
		grpc.CloseFoodGRPCClient()

		os.Exit(0)
	}()
	// Keep the main function running
	select {}
}