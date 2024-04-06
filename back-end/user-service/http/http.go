package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/auth"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/userdb"
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

	userGroup.Post("/api/signup", userdb.CreateUserHandler)
	userGroup.Post("/api/login", auth.LoginUserHandler)
	userGroup.Get("/api/qr/login", auth.LoginQRCodeHandler)
	userGroup.Post("/api/complete", userdb.CompleteUserHandler)
	userGroup.Post("/api/password", userdb.ChangePasswordUserHandler)
	userGroup.Get("/api/qr/recreate", userdb.RecreateQRCodeLogin)
	userGroup.Get("/api/user", userdb.GetUserInfo)

	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
