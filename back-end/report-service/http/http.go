package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/report-service/http/handlers/reports"
)

func StartServer() {
	app := fiber.New()

	// Create a group for routes starting with "/user"
	userGroup := app.Group("/user")

	// Handler for the root endpoint
	userGroup.Get("/", func(c *fiber.Ctx) error {
		// Set the content type header
		c.Set("Content-Type", "text/plain")
		// Return the response string
		return c.SendString("Welcome to the User Service!")
	})

	userApis := userGroup.Group("/api")
	userApis.Get("/users", reports.GetAllUsers)
	userApis.Get("/orders", reports.GetAllOrders)
	userApis.Post("/orders", reports.GetAllOrdersByTime)

	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8010"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
