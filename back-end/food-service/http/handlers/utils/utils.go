package utils

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	card_proto "github.com/nimbolism/software-restaurant/back-end/card-service/proto"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/food-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/gutils"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"gorm.io/gorm"
)

type MealWithCreatedAt struct {
	models.Meal
	CreatedAt time.Time
}

func CheckAccessLevel(AccessLevel int32) bool {
	return AccessLevel > 1
}

func HandleRequest(c *fiber.Ctx, model interface{}, modelName string, createFunc func(interface{}) error) error {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	card, err := GetCardInfo(context.Background(), cookie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Could not get card info: %v", err))
	}

	if !CheckAccessLevel(card.AccessLevel) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "you are not permitted"})
	}

	if err := c.BodyParser(model); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Failed to decode request body: %v", err)})
	}

	db := postgresapp.DB

	// Check if the model already exists
	existingModel := reflect.New(reflect.TypeOf(model).Elem()).Interface()
	if err := db.Where("name = ?", modelName).First(existingModel).Error; err == nil {
		// Model is already taken, return an error
		return fmt.Errorf("%s with this name already exists", modelName)
	}

	if err := createFunc(model); err != nil {
		return fmt.Errorf("failed to insert %s into database: %v", modelName, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": fmt.Sprintf("created %s record", modelName)})
}

func GetCardInfo(ctx context.Context, jwtToken string) (*card_proto.CardInfoResponse, error) {
	if err := grpc.InitializeGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	return grpc.CardServiceClient.GetCardInfo(ctx, &card_proto.GetCardInfoRequest{JwtToken: jwtToken})
}

func (m MealWithCreatedAt) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func GetModels(c *fiber.Ctx, modelName string, modelSlice interface{}) error {
	db := postgresapp.DB
	if err := db.Find(modelSlice).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("%s not found", modelName)})
	}

	// Extract creation time from gorm.Model
	sliceValue := reflect.Indirect(reflect.ValueOf(modelSlice))
	for i := 0; i < sliceValue.Len(); i++ {
		model := sliceValue.Index(i).Interface()
		if modelWithTime, ok := model.(interface{ GetCreatedAt() time.Time }); ok {
			sliceValue.Index(i).FieldByName("CreatedAt").Set(reflect.ValueOf(modelWithTime.GetCreatedAt()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{modelName: modelSlice})
}

func GetPaginatedData(c *fiber.Ctx, db *gorm.DB, model interface{}, perPage int) (interface{}, int64, error) {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	offset := (page - 1) * perPage
	var total int64

	// Get the total count
	if err := db.Model(model).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Create a slice to hold the data
	var data interface{}

	// Paginate the data
	if err := db.Offset(offset).Limit(perPage).Find(model).Error; err != nil {
		return nil, 0, err
	}

	// Assign the data to the interface
	data = model

	return data, total, nil
}

func GenericHandler(c *fiber.Ctx, model interface{}, modelName string, requiredFields []string, errorMessage string) error {
	if err := c.BodyParser(model); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	// Check if required fields are provided
	for _, field := range requiredFields {
		value := reflect.ValueOf(model).Elem().FieldByName(field).String()
		if value == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errorMessage})
		}
	}

	return HandleRequest(c, model, modelName, func(model interface{}) error {
		return postgresapp.DB.Create(model).Error
	})
}

func DeleteRecord(c *fiber.Ctx, modelName string, queryField string, model interface{}) error {
	var request struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	if request.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	db := postgresapp.DB.Where(queryField+" = ?", request.Name)
	if err := db.Delete(model).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete " + modelName})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": modelName + " deleted successfully"})
}
