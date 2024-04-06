package utils

import (
	"bytes"
	"image"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/database"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
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
