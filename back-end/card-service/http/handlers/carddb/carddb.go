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
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
)

func ProfileHandler(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("failed to autherize the user: %v", err)})
	}

	if err := utils.HandleImageUpload(c, authenticateUserResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("failed to upload image: %v", err)})
	}

	newCard := models.Card{
		UserID:      uint(authenticateUserResponse.UserId),
		Reserves:    0,
		BlackListed: false,
		Verified:    false,
		AccessLevel: 1,
	}
	db := postgresapp.DB
	if err := db.Create(&newCard).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Failed to create card: %v", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Image uploaded successfully"})
}

func UpdateImageHandler(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("failed to autherize the user: %v", err)})
	}

	if err := utils.HandleImageUpload(c, authenticateUserResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Failed to uploade image: %v", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Image updated successfully"})
}


// a function to send the image to user
func GetImageHandler(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("failed to autherize the user: %v", err)})
	}

	storageDir := "./uploads"

	imagePath := filepath.Join(storageDir, fmt.Sprintf("%d.jpg", authenticateUserResponse.UserId))

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Image not found"})
	}

	return c.SendFile(imagePath)
}

func GetCardHandler(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("failed to autherize the user: %v", err)})
	}

	card, err := utils.FindCardByUserID(uint(authenticateUserResponse.UserId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve card information"})
	}
	response := map[string]interface{}{
		"reserves":     card.Reserves,
		"blacklisted":  card.BlackListed,
		"verified":     card.Verified,
		"access_level": card.AccessLevel,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func GiveAccessLevel(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("failed to autherize the user: %v", err)})
	}

	// Retrieve card information for the authenticated user
	card, err := utils.FindCardByUserID(uint(authenticateUserResponse.UserId))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to retrieve card information"})
	}

	userAccessLevel := card.AccessLevel
	if err := grpc.InitializeUserGRPCClient(); err != nil {
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Failed in requesting from user service: %v", err)})
	}

	reqCard, err := utils.FindCardByUserID(uint(req.UserId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve requested card information"})
	}

	if userAccessLevel >= reqCard.AccessLevel { //change this and remove equal
		reqCard.AccessLevel++
		db := postgresapp.DB
		err = db.Save(reqCard).Error
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to update card information"})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "successfully updated card access level"})
}

func VerifyUser(c *fiber.Ctx) error {
	authenticateUserResponse, err := grpc.AuthenticateUser(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("failed to autherize the user: %v", err)})
	}

	// Retrieve card information for the authenticated user
	card, err := utils.FindCardByUserID(uint(authenticateUserResponse.UserId))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to retrieve card information"})
	}

	userAccessLevel := card.AccessLevel
	if err := grpc.InitializeUserGRPCClient(); err != nil {
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Failed in requesting from user service: %v", err)})
	}

	reqCard, err := utils.FindCardByUserID(uint(req.UserId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve requested card information"})
	}

	if userAccessLevel >= 2 { 
		reqCard.Verified = true
		db := postgresapp.DB
		err = db.Save(reqCard).Error
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to update card information"})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "successfully updated card verification status"})
}
