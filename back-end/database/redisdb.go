package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// RedisConfig holds the configuration for the Redis connection.
type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

// NewRedisConfig creates a new RedisConfig instance from environment variables.
func NewRedisConfig() (*RedisConfig, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	if host == "" || port == "" {
		return nil, fmt.Errorf("one or more Redis environment variables are not set")
	}

	return &RedisConfig{
		Host:     host,
		Port:     port,
		Password: password,
	}, nil
}

// InitRedisDB initializes the connection to the Redis database.
func InitRedisDB(ctx context.Context, cfg *RedisConfig) (*redis.Client, error) {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
	}

	client := redis.NewClient(options)

	// Test the Redis database connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to the Redis database: %v", err)
	}

	log.Println("Connected to the Redis database successfully!")
	return client, nil
}

// CloseRedisDB closes the Redis database connection.
func CloseRedisDB(client *redis.Client) error {
	if client != nil {
		if err := client.Close(); err != nil {
			return fmt.Errorf("failed to close Redis connection: %v", err)
		}
		log.Println("Closed Redis connection successfully!")
	}
	return nil
}
