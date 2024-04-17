package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/auth"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/userdb"
)

func StartServer() {
	app := fiber.New()

	// /user routes group
	userGroup := app.Group("/user")

	// setting root endpoint for testing
	userGroup.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendString("Welcome to the User Service!")
	})

	// /api for api routes group
	userApis := userGroup.Group("/api")
	userApis.Post("/signup", userdb.CreateUserHandler)
	userApis.Post("/login", auth.LoginUserHandler)
	userApis.Get("/qr/login", auth.LoginQRCodeHandler)
	userApis.Put("/complete", userdb.CompleteUserHandler)
	userApis.Post("/password", userdb.ChangePasswordUserHandler)
	userApis.Get("/qr/recreate", userdb.RecreateQRCodeLogin)
	userApis.Get("/user", userdb.GetUserInfo)

	// starting the http server
	log.Println("Starting user HTTP server...")
	if err := app.Listen(":8010"); err != nil {
		log.Fatalf("Failed to start user HTTP server: %v", err)
	}
}
