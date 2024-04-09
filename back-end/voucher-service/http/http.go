package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New()

	// Create a group for routes starting with "/user"
	voucherGroup := app.Group("/voucher")

	// Handler for the root endpoint
	voucherGroup.Get("/", func(c *fiber.Ctx) error {
		// Set the content type header
		c.Set("Content-Type", "text/plain")
		// Return the response string
		return c.SendString("Welcome to the voucher Service!")
	})

	// orderApis := orderGroup.Group("/api")

	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8084"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
