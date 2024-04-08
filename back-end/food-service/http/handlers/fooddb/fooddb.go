package fooddb

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/food-service/http/handlers/utils"
)

func FoodHandler(c *fiber.Ctx) error {
	var foodRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Category    string `json:"category"` // Change the type to string
		Meal        string `json:"meal"`
	}

	if err := c.BodyParser(&foodRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	// Check if required fields are provided
	if foodRequest.Name == "" || foodRequest.Description == "" || foodRequest.Category == "" || foodRequest.Meal == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name, Description, Category, and Meal are required"})
	}

	// Fetch CategoryID based on the provided name
	var category models.Category
	if err := database.GetPQDB().Where("name = ?", foodRequest.Category).First(&category).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category not found"})
	}

	var meal models.Meal
	if err := database.GetPQDB().Where("name = ?", foodRequest.Meal).First(&meal).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "meal not found"})
	}

	// Create Food object
	newFood := models.Food{
		Name:        foodRequest.Name,
		Description: foodRequest.Description,
		CategoryID:  category.ID,
		MealID:      meal.ID,
	}

	return utils.HandleRequest(c, &newFood, "Food", func(model interface{}) error {
		return database.GetPQDB().Create(model).Error
	})
}

func CategoryHandler(c *fiber.Ctx) error {
	var category models.Category
	return utils.HandleRequest(c, &category, "Category", func(model interface{}) error {
		return database.GetPQDB().Create(model).Error
	})
}

func MealHandler(c *fiber.Ctx) error {
	var meal models.Meal
	return utils.HandleRequest(c, &meal, "Meal", func(model interface{}) error {
		return database.GetPQDB().Create(model).Error
	})
}

func SideDishHandler(c *fiber.Ctx) error {
	var sideDishRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&sideDishRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	// Check if required fields are provided
	if sideDishRequest.Name == "" || sideDishRequest.Description == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name and Description are required"})
	}

	// Create SideDish object
	newSideDish := models.SideDish{
		Name:        sideDishRequest.Name,
		Description: sideDishRequest.Description,
	}

	return utils.HandleRequest(c, &newSideDish, "SideDish", func(model interface{}) error {
		return database.GetPQDB().Create(model).Error
	})
}

func GetFoods(c *fiber.Ctx) error {
	perPage, err := strconv.Atoi(c.Query("perPage", "20"))
	if err != nil || perPage < 1 {
		perPage = 20
	}

	db := database.GetPQDB().Model(&models.Food{})

	// Define category and meal names to filter by
	categoryName := c.Query("category", "")
	mealName := c.Query("meal", "")

	// Join the Category and Meal tables and select their fields
	db = db.Table("foods").
		Select("foods.*, categories.name AS category_name, meals.name AS meal_name").
		Joins("LEFT JOIN categories ON foods.category_id = categories.id").
		Joins("LEFT JOIN meals ON foods.meal_id = meals.id")

	// Filter by category name if provided
	if categoryName != "" {
		db = db.Where("categories.name = ?", categoryName)
	}

	// Filter by meal name if provided
	if mealName != "" {
		db = db.Where("meals.name = ?", mealName)
	}

	type Result struct {
		models.Food
		CategoryName string
		MealName     string
	}

	// Get paginated data
	var results []Result
	data, total, err := utils.GetPaginatedData(c, db, &results, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Define pagination variables
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	previous := page > 1
	next := total > int64(page*perPage)

	// Prepare response data
	responseData := fiber.Map{
		"page":        page,
		"perPage":     perPage,
		"total":       total,
		"previous":    previous,
		"next":        next,
		"data":        data,
		"total_pages": totalPages,
	}

	// Return response
	return c.Status(fiber.StatusOK).JSON(responseData)
}

func GetSideDishes(c *fiber.Ctx) error {
	perPage, err := strconv.Atoi(c.Query("perPage", "20"))
	if err != nil || perPage < 1 {
		perPage = 20
	}

	db := database.GetPQDB().Model(&models.SideDish{})

	// Create a slice to hold the data
	var sideDishes []models.SideDish

	// Get paginated data
	data, total, err := utils.GetPaginatedData(c, db, &sideDishes, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	page, err := strconv.Atoi(c.Query("page", "1")) // Define page here
	if err != nil || page < 1 {
		page = 1
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))
	previous := page > 1
	next := total > int64(page*perPage)

	responseData := fiber.Map{
		"page":        page,
		"perPage":     perPage,
		"total":       total,
		"previous":    previous,
		"next":        next,
		"data":        data,
		"total_pages": totalPages,
	}

	return c.Status(fiber.StatusOK).JSON(responseData)
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	return utils.GetModels(c, "Categories", &categories)
}

func GetMeals(c *fiber.Ctx) error {
	var meals []models.Meal
	return utils.GetModels(c, "Meals", &meals)
}
