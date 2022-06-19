package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Returns an error response
func RespondWithError(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"message": strings.ToLower(err.Error()),
	})
}
