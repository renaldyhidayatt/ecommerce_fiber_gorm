package handler

import (
	"ecommerce_fiber/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) initSecretGroup(api *fiber.App) {

	secret := api.Group("/secret")

	api.Use(middleware.Protector())

	secret.Get("/", h.RestrictedEndpoint)
}

func (h *Handler) RestrictedEndpoint(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["sub"].(string)
	return c.SendString("Welcome " + name)
}
