package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	// Create a group for routes starting with "/user"
	orderGroup := app.Group("/order")

	// Handler for the root endpoint
	orderGroup.Get("/", func(c *fiber.Ctx) error {
		// Set the content type header
		c.Set("Content-Type", "text/plain")
		// Return the response string
		return c.SendString("Welcome to the order Service!")
	})

	// orderApis := orderGroup.Group("/api")

	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8083"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
