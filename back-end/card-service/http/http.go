package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/card-service/http/handlers/carddb"
)

func StartServer() {
	app := fiber.New()

	// /card group for card-service
	cardGroup := app.Group("/card")

	// root handler for testing
	cardGroup.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendString("Welcome to the Card Service!")
	})

	// /api group for APIs
	cardApis := cardGroup.Group("/api")
	cardApis.Post("/photo", carddb.ProfileHandler)
	cardApis.Post("/photo/new", carddb.UpdateImageHandler)
	cardApis.Post("/access", carddb.GiveAccessLevel)
	cardApis.Post("/verify", carddb.VerifyUser)
	cardApis.Get("/photo", carddb.GetImageHandler)
	cardApis.Get("/card", carddb.GetCardHandler)

	// Start card HTTP server
	println("Starting card HTTP server...")
	if err := app.Listen(":8020"); err != nil {
		log.Fatalf("Failed to start card HTTP server: %v", err)
	}
}
