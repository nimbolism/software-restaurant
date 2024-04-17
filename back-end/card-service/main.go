package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nimbolism/software-restaurant/back-end/card-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/card-service/http"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
)

func main() {
	postgresapp := postgresapp.New()
	defer postgresapp.Close()

	if err := grpc.InitializeUserGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	go grpc.StartServer()

	go http.StartServer()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		// Close the user gRPC client connection when the server is shutting down
		grpc.CloseGRPCClient()

		os.Exit(0)
	}()

	// Keep the main function running
	select {}
}
