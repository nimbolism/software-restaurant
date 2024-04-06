package carddb

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/card-service/grpc"
	"github.com/nimbolism/software-restaurant/back-end/card-service/http/handlers/utils"
	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/database/models"

	"github.com/nimbolism/software-restaurant/back-end/gutils"
)

func ProfileHandler(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	if err := grpc.InitializeGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	// Call the AuthenticateUser function of the user-service
	authenticateUserResponse, err := grpc.AuthenticateUser(context.Background(), cookie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to authenticate user via gRPC")
	}

	// If authentication fails, populate the error field and return
	if !authenticateUserResponse.Success {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": authenticateUserResponse.ErrorMessage})
	}

	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Unable to parse JSON body")
	}

	// Get the Base64 image from the JSON body
	base64Image, ok := body["image"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("No image in JSON body")
	}

	base64Image = strings.SplitN(base64Image, ",", 2)[1]
	// Decode the Base64 image
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Unable to decode Base64 image")
	}

	// Compress the image to 50KB
	compressedImageData, err := utils.CompressImage(imageData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to compress image")
	}

	// Generate a secure filename
	filename := fmt.Sprintf("%d.jpg", authenticateUserResponse.UserId)

	// Specify the directory where images will be stored (make sure this directory exists)
	storageDir := "./uploads"
	err = os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create storage directory")
	}

	// Save the image to a secure file
	err = os.WriteFile(filepath.Join(storageDir, filename), compressedImageData, 0644)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save image to file")
	}

	// Create a new card record
	newCard := models.Card{
		UserID:      uint(authenticateUserResponse.UserId),
		Reserves:    0,
		BlackListed: false,
		AccessLevel: 1,
	}
	db := database.GetPQDB()
	if err := db.Create(&newCard).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to create card: %v", err))
	}

	return c.SendString("Image uploaded successfully")
}

func UpdateImageHandler(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	if err := grpc.InitializeGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	// Call the AuthenticateUser function of the user-service
	authenticateUserResponse, err := grpc.AuthenticateUser(context.Background(), cookie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to authenticate user via gRPC")
	}

	// If authentication fails, populate the error field and return
	if !authenticateUserResponse.Success {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": authenticateUserResponse.ErrorMessage})
	}

	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Unable to parse JSON body")
	}

	// Get the Base64 image from the JSON body
	base64Image, ok := body["image"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("No image in JSON body")
	}

	base64Image = strings.SplitN(base64Image, ",", 2)[1]
	// Decode the Base64 image
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Unable to decode Base64 image")
	}

	// Compress the image to 50KB
	compressedImageData, err := utils.CompressImage(imageData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to compress image")
	}

	// Generate a secure filename
	filename := fmt.Sprintf("%d.jpg", authenticateUserResponse.UserId)

	// Specify the directory where images will be stored (make sure this directory exists)
	storageDir := "./uploads"

	// Save the updated image to a secure file, overwriting the existing image
	err = os.WriteFile(filepath.Join(storageDir, filename), compressedImageData, 0644)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update image")
	}

	return c.SendString("Image updated successfully")
}


func GetImageHandler(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	if err := grpc.InitializeGRPCClient(); err != nil {
		log.Fatalf("Failed to initialize gRPC client: %v", err)
	}

	// Call the AuthenticateUser function of the user-service
	authenticateUserResponse, err := grpc.AuthenticateUser(context.Background(), cookie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to authenticate user via gRPC: %v", err))
	}

	// If authentication fails, populate the error field and return
	if !authenticateUserResponse.Success {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized access"})
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
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	// Authenticate user via gRPC (you can reuse the existing function)
	authenticateUserResponse, err := grpc.AuthenticateUser(context.Background(), cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": fmt.Sprintf("Failed to authenticate user via gRPC: %v", err)})
	}

	// Retrieve card information for the authenticated user
	card, err := utils.FindCardByUserID(uint(authenticateUserResponse.UserId))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to retrieve card information"})
	}
	response := map[string]interface{}{
		"reserves":     card.Reserves,
		"blacklisted":  card.BlackListed,
		"access_level": card.AccessLevel,
	}

	return c.JSON(response)
}
