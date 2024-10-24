package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msvc_rol/infrastructure/utils"
)

func ValidateToken(c *fiber.Ctx) error {
	// Validate JWT token here

	token := c.Get(utils.AUTHORIZATION)

	if len(token) > 7 && token[:7] == utils.BEARER {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		utils.STATUS:  fiber.StatusUnauthorized,
		utils.MESSAGE: utils.TOKEN_INVALID,
	})

}
