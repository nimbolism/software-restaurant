package carddb

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/card-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/card-service/http/handlers/utils"
	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
)

func ProfileHandler(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return err
	}

	if err := utils.HandleImageUpload(c, authenticateUserResponse); err != nil {
		return err
	}

	// Create a new card record
	newCard := models.Card{
		UserID:      uint(authenticateUserResponse.UserId),
		Reserves:    0,
		BlackListed: false,
		Verified:    false,
		AccessLevel: 1,
	}
	db := database.GetPQDB()
	if err := db.Create(&newCard).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to create card: %v", err))
	}

	return c.SendString("Image uploaded successfully")
}

func UpdateImageHandler(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return err
	}

	if err := utils.HandleImageUpload(c, authenticateUserResponse); err != nil {
		return err
	}

	return c.SendString("Image updated successfully")
}

func GetImageHandler(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return err
	}

	// Specify the directory where images are stored
	storageDir := "./uploads"

	// Create the full path of the image
	imagePath := filepath.Join(storageDir, fmt.Sprintf("%d.jpg", authenticateUserResponse.UserId))

	// Check if the image exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("Image not found")
	}

	// Send the image to the user
	return c.SendFile(imagePath)
}

func GetCardHandler(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return err
	}

	// Retrieve card information for the authenticated user
	card, err := utils.FindCardByUserID(uint(authenticateUserResponse.UserId))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to retrieve card information"})
	}
	response := map[string]interface{}{
		"reserves":     card.Reserves,
		"blacklisted":  card.BlackListed,
		"verified":     card.Verified,
		"access_level": card.AccessLevel,
	}

	return c.JSON(response)
}

func GiveAccessLevel(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return err
	}

	// Retrieve card information for the authenticated user
	card, err := utils.FindCardByUserID(uint(authenticateUserResponse.UserId))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to retrieve card information"})
	}

	userAccessLevel := card.AccessLevel
	if err := grpc.InitializeGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	var reqBody struct {
		Username string `json:"username"`
	}
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	req, err := grpc.UserServiceClient.GetOneUser(context.Background(), &user_proto.GetOneUserRequest{Username: reqBody.Username})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Failed in requesting from user service: %v, service: %v", err.Error(), req.ErrorMessage)})
	}

	reqCard, err := utils.FindCardByUserID(uint(req.UserId))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to retrieve requested card information"})
	}

	if userAccessLevel >= reqCard.AccessLevel { //change this and remove equal
		reqCard.AccessLevel++
		db := database.GetPQDB()
		err = db.Save(reqCard).Error
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to update card information"})
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": "successfully updated card access level"})
}
