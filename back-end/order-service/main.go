package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/order-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/order-service/http"
)

func main() {
	postgresapp := postgresapp.New()
	defer postgresapp.Close()

	if err := grpc.InitializeUserGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	if err := grpc.InitializeCardGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	if err := grpc.InitializeFoodGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	if err := grpc.InitializeVoucherGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}
	
	go grpc.StartServer()

	go http.StartServer()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		grpc.CloseUserGRPCClient()
		grpc.CloseCardGRPCClient()
		grpc.CloseFoodGRPCClient()
		grpc.CloseVoucherGRPCClient()

		os.Exit(0)
	}()
	
	// Keep the main function running
	select {}
}
