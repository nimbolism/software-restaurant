package gutils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookie(c *fiber.Ctx, name, value string, expiry time.Time) error {
	// Set the cookie with the provided parameters
	c.Cookie(&fiber.Cookie{
		Name:    name,
		Value:   value,
		Expires: expiry,
		// Secure:   true,
		// HTTPOnly: true,
		// SameSite: "None",
	})

	return nil
}

func GetCookie(c *fiber.Ctx, name string) string {
	return c.Cookies(name)
}
