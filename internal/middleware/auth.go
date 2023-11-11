package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Protector() fiber.Handler {
	signingKey := []byte(viper.GetString("JWT_SECRET"))
	config := jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: signingKey},
		ErrorHandler: jwtError,
	}
	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {

	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
