package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/report-service/http/handlers/reports"
)

func StartServer() {
	app := fiber.New()

	// Create a group for routes starting with "/report"
	reportGroup := app.Group("/report")

	// Handler for the root endpoint
	reportGroup.Get("/", func(c *fiber.Ctx) error {
		// Set the content type header
		c.Set("Content-Type", "text/plain")
		// Return the response string
		return c.SendString("Welcome to the Report Service!")
	})

	reportApis := reportGroup.Group("/api")
	reportApis.Get("/users", reports.GetAllUsers)
	reportApis.Get("/orders", reports.GetAllOrders)
	reportApis.Post("/orders", reports.GetAllOrdersByTime)
	reportApis.Get("/vouchers", reports.GetAllOrdersVoucher)

	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8060"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
