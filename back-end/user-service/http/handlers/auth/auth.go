package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/gutils"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/utils"
	"github.com/skip2/go-qrcode"
)

func Login(username, password string) (string, error) {
	var user models.User
	db := postgresapp.DB

	// Find user by username
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return "", fmt.Errorf("failed to find user: %v", result.Error)
	}

	// Compare hashed password with the provided password
	err := utils.ComparePasswords(user.Password, password)
	if err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	// Generate JWT token
	expiry := time.Now().Add(24 * time.Hour)
	token, err := GenerateJWT(username, expiry)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %v", err)
	}

	// Return JWT token on successful login
	return token, nil
}

// LoginUserHandler handles HTTP requests to login a user
func LoginUserHandler(c *fiber.Ctx) error {
	var loginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	// Login user
	token, err := Login(loginUser.Username, loginUser.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	// Set JWT token as a cookie
	expiry := time.Now().Add(24 * time.Hour)
	if err := gutils.SetCookie(c, "jwt", token, expiry); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to set JWT token as secure cookie"})
	}

	// Respond with success message
	return c.SendString("Login successful")
}

func GenerateJWT(username string, expiry time.Time) (string, error) {
	// Retrieve the secret key from environment variables
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("secret key not found in environment variables")
	}

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = expiry.Unix() // Token expires based on the provided duration

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %v", err)
	}

	return tokenString, nil
}

func GetUsernameFromJWT(cookie string) (string, error) {
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse JWT token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid JWT token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username not found in JWT claims")
	}

	return username, nil
}

// CreateQRCode generates a QR code for a user
func CreateQRCode(username string) ([]byte, error) {
	// Generate a JWT for the user
	expiry := time.Now().Add(72 * time.Hour)
	token, err := GenerateJWT(username, expiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token: %v", err)
	}

	// Generate QR code for the user
	qrCode, err := qrcode.New(fmt.Sprintf("https://localhost:5050/user/api/qr/login?token=%s", token), qrcode.Medium)
	if err != nil {
		return nil, fmt.Errorf("failed to create QR code: %v", err)
	}

	// Encode the QR code to PNG
	png, err := qrCode.PNG(256)
	if err != nil {
		return nil, fmt.Errorf("failed to encode QR code to PNG: %v", err)
	}

	return png, nil
}

// LoginQRCodeHandler handles HTTP requests to verify QR code and set JWT token cookie for the user
func LoginQRCodeHandler(c *fiber.Ctx) error {
	// Extract token from URL parameter
	token := c.Query("token")

	// Verify the token
	username, err := GetUsernameFromJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	// Generate JWT token
	expiry := time.Now().Add(24 * time.Hour)
	jwttoken, err := GenerateJWT(username, expiry)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate JWT token"})
	}
	if err := gutils.SetCookie(c, "jwt", jwttoken, expiry); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to set JWT token as secure cookie"})
	}

	// Respond with success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "user is logged in"})
}
