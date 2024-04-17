package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nimbolism/software-restaurant/back-end/database/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashedPassword), nil
}

func ComparePasswords(firstPassword, secondpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(firstPassword), []byte(secondpassword))
}

func GetUserFromContext(c *fiber.Ctx) (models.User, error) {
	var user models.User
	err := c.BodyParser(&user)
	return user, err
}

func GetExistingUser(db *gorm.DB, username string) (models.User, error) {
	var user models.User
	err := db.Where("username = ?", username).First(&user).Error
	return user, err
}

func IsUnique(db *gorm.DB, field string, value string, username string) bool {
	var count int64
	db.Model(&models.User{}).Where(field+" = ?", value).Not("username = ?", username).Count(&count)
	return count == 0
}

func UpdateUserData(existingUser *models.User, completingUser models.User) {
	existingUser.Email = completingUser.Email
	existingUser.PhoneNumber = completingUser.PhoneNumber
	existingUser.NationalCode = completingUser.NationalCode
}

func SaveUser(db *gorm.DB, user *models.User) error {
	return db.Save(user).Error
}
