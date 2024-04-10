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

	userApis := userGroup.Group("/api")
	userApis.Post("/signup", userdb.CreateUserHandler)
	userApis.Post("/login", auth.LoginUserHandler)
	userApis.Get("/qr/login", auth.LoginQRCodeHandler)
	userApis.Put("/complete", userdb.CompleteUserHandler)
	userApis.Post("/password", userdb.ChangePasswordUserHandler)
	userApis.Get("/qr/recreate", userdb.RecreateQRCodeLogin)
	userApis.Get("/user", userdb.GetUserInfo)

	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8010"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
