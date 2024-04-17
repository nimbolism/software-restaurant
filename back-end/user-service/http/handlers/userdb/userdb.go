package userdb

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"github.com/nimbolism/software-restaurant/back-end/gutils"
	"github.com/nimbolism/software-restaurant/back-end/gutils/postgresapp"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/auth"
	"github.com/nimbolism/software-restaurant/back-end/user-service/http/handlers/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreateUser(username, password, confirmPassword string) error {
	// Check if password and confirm_password match
	if password != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	db := postgresapp.DB
	// Check if the username already exists
	var existingUser models.User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		// Username is already taken, return an error
		return fmt.Errorf("username already exists: %v", err)
	}

	qr, err := auth.CreateQRCode(username)
	if err != nil {
		return fmt.Errorf("failed to create qr code: %v", err)
	}

	// Create the user
	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		QRCode:   qr,
	}
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to insert user into database: %v", err)
	}

	return nil
}

func CreateUserHandler(c *fiber.Ctx) error {
	var signUpUser struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	if err := c.BodyParser(&signUpUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	// Create user
	if err := CreateUser(signUpUser.Username, signUpUser.Password, signUpUser.ConfirmPassword); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "User created successfully"})
}

func CompleteUserHandler(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	username, err := auth.GetUsernameFromJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	completingUser, err := utils.GetUserFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	db := postgresapp.DB
	existingUser, err := utils.GetExistingUser(db, username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Check for uniqueness
	fields := []string{"email", "phone_number", "national_code"}
	for _, field := range fields {
		value := reflect.ValueOf(completingUser).FieldByName(cases.Title(language.English).String(field)).String()
		if value != "" && !utils.IsUnique(db, field, value, username) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": field + " is already taken"})
		}
	}

	utils.UpdateUserData(&existingUser, completingUser)

	if err := utils.SaveUser(db, &existingUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to complete user data"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "User data completed successfully"})
}

func ChangePasswordUserHandler(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	username, err := auth.GetUsernameFromJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	var changePasswordData struct {
		OldPassword     string `json:"old_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	if err := c.BodyParser(&changePasswordData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to decode request body"})
	}

	// Retrieve the user from the database
	db := postgresapp.DB
	var existingUser models.User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Compare old password
	if err := utils.ComparePasswords(existingUser.Password, changePasswordData.OldPassword); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect old password"})
	}

	// Check if new password matches confirm password
	if changePasswordData.NewPassword != changePasswordData.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "New password and confirm password do not match"})
	}

	// Hash the new password
	hashedNewPassword, err := utils.HashPassword(changePasswordData.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash new password"})
	}

	// Update user's password
	existingUser.Password = string(hashedNewPassword)
	if err := db.Save(&existingUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update password"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "Password changed successfully"})
}

func RecreateQRCodeLogin(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	username, err := auth.GetUsernameFromJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	qrcode, err := auth.CreateQRCode(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create qrcode!"})
	}

	db := postgresapp.DB
	existingUser, err := utils.GetExistingUser(db, username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	// update qrcode
	existingUser.QRCode = qrcode

	if err := utils.SaveUser(db, &existingUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to complete user data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "updated user's qrcode"})
}

func GetUserInfo(c *fiber.Ctx) error {
	cookie := gutils.GetCookie(c, "jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "JWT cookie not found"})
	}

	username, err := auth.GetUsernameFromJWT(cookie)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	db := postgresapp.DB
	existingUser, err := utils.GetExistingUser(db, username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	userInfo := fiber.Map{
		"username":      existingUser.Username,
		"email":         existingUser.Email,
		"qr_code":       existingUser.QRCode,
		"phone_number":  existingUser.PhoneNumber,
		"national_code": existingUser.NationalCode,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": userInfo})
}
