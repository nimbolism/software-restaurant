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
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	user_proto "github.com/nimbolism/software-restaurant/back-end/user-service/proto"
)

func FindCardByUserID(userID uint) (*models.Card, error) {
	var card models.Card
	db := postgresapp.DB
	if err := db.Where("user_id = ?", userID).First(&card).Error; err != nil {
		return nil, err
	}
	return &card, nil
}

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

	// A loop to reduce image quality to 50KB or less
	for {
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unable to parse JSON body"})
	}

	// Get the image from body
	base64Image, ok := body["image"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No image in JSON body"})
	}

	// remove metadata of base64 image & decode it
	base64Image = strings.SplitN(base64Image, ",", 2)[1]
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to decode Base64 image"})
	}

	// Compress the image to 50KB
	compressedImageData, err := CompressImage(imageData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to compress image"})
	}

	// Generate a secure filename using user_ID
	filename := fmt.Sprintf("%d.jpg", authenticateUserResponse.UserId)

	storageDir := "./uploads"
	err = os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create storage directory"})
	}

	// Saving the image
	err = os.WriteFile(filepath.Join(storageDir, filename), compressedImageData, 0644)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save image to file"})
	}

	return nil
}
