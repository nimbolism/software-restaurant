package main

import (
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/user-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http"
)

func main() {
	postgresapp := postgresapp.New()
	defer postgresapp.Close()

	go grpc.StartServer()

	go http.StartServer()

	// Keep the main function running
	select {}
}
