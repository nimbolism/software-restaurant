// http/server.go
package http

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/user-service/database/userdb"
)

func StartServer(db *sql.DB) {
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

	userGroup.Post("/api/user", func(c *fiber.Ctx) error {
		// Parse request body
		var user userdb.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
		}

		// Add user to the database
		if err := userdb.AddUser(&user, db); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add user to the database"})
		}

		// Return the created user ID in the response
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"userId": user.UserID})
	})

	userGroup.Get("/api/user/:userId", func(c *fiber.Ctx) error {
		// Extract user ID from URL path parameters
		userID := c.Params("userId")

		// Retrieve user information from the database
		user, err := userdb.GetUserByID(userID, db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user information"})
		}

		// Return user information in the response
		return c.JSON(user)
	})

	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
