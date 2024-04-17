package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/report-service/http/handlers/reports"
)

func StartServer() {
	app := fiber.New()

	// /report group for report service
	reportGroup := app.Group("/report")

	// root endpoint for testing
	reportGroup.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendString("Welcome to the Report Service!")
	})

	// /api for apis of report service
	reportApis := reportGroup.Group("/api")
	reportApis.Get("/users", reports.GetAllUsers)
	reportApis.Get("/orders", reports.GetAllOrders)
	reportApis.Post("/orders", reports.GetAllOrdersByTime)
	reportApis.Get("/vouchers", reports.GetAllOrdersVoucher)

	println("Starting report HTTP server...")
	if err := app.Listen(":8060"); err != nil {
		log.Fatalf("Failed to start report HTTP server: %v", err)
	}
}
