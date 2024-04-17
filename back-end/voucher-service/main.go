package main

import (
	"context"
	"log"

	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/voucher-service/grpc"
)

func main() {
	redisCfg, err := database.NewRedisConfig()
	if err != nil {
		log.Fatalf("Failed to create Redis configuration: %v", err)
	}
	redisClient, err := database.InitRedisDB(context.Background(), redisCfg)
	if err != nil {
		log.Fatalf("Failed to initialize Redis database: %v", err)
	}
	defer func() {
		if err := database.CloseRedisDB(redisClient); err != nil {
			log.Fatalf("Error closing Redis connection: %v", err)
		}
	}()

	go grpc.StartServer()

	// Keep the main function running
	select {}
}
