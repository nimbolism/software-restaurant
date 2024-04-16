package main

import "github.com/nimbolism/software-restaurant/back-end/report-service/http"

func main() {
	// Start gRPC server
	go http.StartServer()

	// Keep the main function running
	select {}
}
