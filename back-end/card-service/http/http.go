package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/card-service/http/handlers/carddb"
)

func StartServer() {
	app := fiber.New()

	// Create a group for routes starting with "/user"
	userGroup := app.Group("/card")

	// Handler for the root endpoint
	userGroup.Get("/", func(c *fiber.Ctx) error {
		// Set the content type header
		c.Set("Content-Type", "text/plain")
		// Return the response string
		return c.SendString("Welcome to the Card Service!")
	})
	userGroup.Post("/api/photo", carddb.ProfileHandler)
	userGroup.Post("/api/photo/new", carddb.UpdateImageHandler)
	userGroup.Get("/api/photo", carddb.GetImageHandler)
	userGroup.Get("/api/card", carddb.GetCardHandler)

	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8081"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
