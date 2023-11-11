package handler

import (
	"ecommerce_fiber/internal/domain/requests/auth"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initAuthGroup(api *fiber.App) {
	auth := api.Group("/api/auth")

	auth.Get("/hello", h.handlerHello)
	auth.Post("/register", h.register)
	auth.Post("/login", h.login)
}

func (h *Handler) handlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Auth")
}

func (h *Handler) register(c *fiber.Ctx) error {
	var body auth.RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})

	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}

	res, err := h.services.User.Register(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":      false,
		"message":    "Register new user account success",
		"data":       res,
		"statusCode": fiber.StatusOK,
	})
}

func (h *Handler) login(c *fiber.Ctx) error {
	var body auth.LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res, err := h.services.User.Login(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login susccses",
		"data":    res,
	})

}
