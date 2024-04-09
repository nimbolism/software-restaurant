package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/food-service/http/handlers/fooddb"
)

func StartServer() {
	app := fiber.New()

	// Create a group for routes starting with "/user"
	foodGroup := app.Group("/food")

	// Handler for the root endpoint
	foodGroup.Get("/", func(c *fiber.Ctx) error {
		// Set the content type header
		c.Set("Content-Type", "text/plain")
		// Return the response string
		return c.SendString("Welcome to the Food Service!")
	})

	foodApis := foodGroup.Group("/api")
	foodApis.Post("/new/food", fooddb.FoodHandler)
	foodApis.Post("/new/category", fooddb.CategoryHandler)
	foodApis.Post("/new/meal", fooddb.MealHandler)
	foodApis.Post("/new/sidedish", fooddb.SideDishHandler)
	
	foodApis.Delete("/food", fooddb.DelFood)
	foodApis.Delete("/category", fooddb.DelCategory)
	foodApis.Delete("/meal", fooddb.DelMeal)
	foodApis.Delete("/sidedish", fooddb.DelSideDish)

	foodApis.Get("/food", fooddb.GetFoods)
	foodApis.Get("/category", fooddb.GetCategories)
	foodApis.Get("/meal", fooddb.GetMeals)
	foodApis.Get("/sidedish", fooddb.GetSideDishes)


	// Start Fiber HTTP server
	println("Starting Fiber HTTP server...")
	if err := app.Listen(":8082"); err != nil {
		log.Fatalf("Failed to start Fiber HTTP server: %v", err)
	}
}
