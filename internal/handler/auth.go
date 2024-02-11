package handler

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initAuthGroup(api *fiber.App) {
	auth := api.Group("/api/auth")

	auth.Get("/hello", h.handlerHello)
	auth.Post("/register", h.register)
	auth.Post("/login", h.login)
	auth.Post("/refresh-token", h.RefreshToken)
}

// handlerHello function
// @Summary Greet the user
// @Description Return a greeting message to the user
// @Tags Auth
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /auth/hello [get]
func (h *Handler) handlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Auth")
}

// register function
// @Summary Register to the application
// @Description Create User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body auth.RegisterRequest true "User information"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /auth/register [post]
func (h *Handler) register(c *fiber.Ctx) error {
	var body auth.RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})

	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Auth.Register(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Register success",
		Data:       res,
	})
}

// register function
// @Summary Login to the application
// @Description Login with email and password to get a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body auth.LoginRequest true "User information"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /auth/login [post]
func (h *Handler) login(c *fiber.Ctx) error {
	var body auth.LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Auth.Login(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Login success",
		Data:       res,
	})

}

// register function
// @Summary Refresh Token to the application
// @Description Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body auth.RefreshTokenRequest true "User information"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /auth/refresh-token [post]
func (h *Handler) RefreshToken(c *fiber.Ctx) error {
	var body auth.RefreshTokenRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Auth.RefreshToken(body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Login success",
		Data:       res,
	})

}
