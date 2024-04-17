package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/food-service/http/handlers/fooddb"
)

func StartServer() {
	app := fiber.New()

	// /food group for food service
	foodGroup := app.Group("/food")

	// root endpointn for testing
	foodGroup.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.SendString("Welcome to the Food Service!")
	})

	// /api group for apis of food service
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


	println("Starting food HTTP server...")
	if err := app.Listen(":8030"); err != nil {
		log.Fatalf("Failed to start food HTTP server: %v", err)
	}
}
