package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
)

func FindCardByUserID(userID uint) (*models.Card, error) {
	var card models.Card
	db := database.GetPQDB()
	if err := db.Where("user_id = ?", userID).First(&card).Error; err != nil {
		return nil, err
	}
	return &card, nil
}

// Function to check if the uploaded file is an image
func IsImage(c *fiber.Ctx, fieldName string) bool {
	file, err := c.FormFile("image")
	if err != nil {
		return false
	}
	contentType := file.Header.Get("Content-Type")
	return strings.HasPrefix(contentType, "image/")
}

func CompressImage(imageData []byte) ([]byte, error) {
	// Decode image
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		return nil, err
	}

	// Initialize quality and dimensions
	quality := 100
	maxDimension := 1280

	for {
		// Resize image
		if img.Bounds().Dx() > maxDimension || img.Bounds().Dy() > maxDimension {
			img = imaging.Resize(img, maxDimension, maxDimension, imaging.Lanczos)
		}

		// Compress image
		var buf bytes.Buffer
		err = imaging.Encode(&buf, img, imaging.JPEG, imaging.JPEGQuality(quality))
		if err != nil {
			return nil, err
		}

		// If the image size is less than 50KB or the quality is too low, stop
		if buf.Len() < 50000 || quality < 10 {
			return buf.Bytes(), nil
		}

		// Reduce quality and dimensions
		quality -= 10
		maxDimension -= 100
	}
}

func HandleImageUpload(c *fiber.Ctx, authenticateUserResponse *user_proto.AuthenticateUserResponse) error {
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
	compressedImageData, err := CompressImage(imageData)
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

	return nil
}
