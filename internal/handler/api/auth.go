package api

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type authHandleApi struct {
	client pb.AuthServiceClient
}

func NewAuthHandleApi(client pb.AuthServiceClient, router *fiber.App) {
	api := &authHandleApi{
		client: client,
	}

	routerAuth := router.Group("/api/auth")

	routerAuth.Get("/hello", api.handlerHello)

	routerAuth.Post("/api/auth/login", api.handleLogin)
	routerAuth.Post("/api/auth/register", api.handleRegister)
	routerAuth.Post("/api/auth/refresh-token", api.handleRefreshToken)
}

// handlerHello function
// @Summary Greet the user
// @Description Return a greeting message to the user
// @Tags Auth
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /auth/hello [get]
func (h *authHandleApi) handlerHello(c *fiber.Ctx) error {
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
func (h *authHandleApi) handleRegister(c *fiber.Ctx) error {
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

	res, err := h.client.RegisterUser(c.Context(), &pb.RegisterRequest{
		Name:            body.Name,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.Confirm_password,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.JSON(response.Response{
		Message:    "Success",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})

}

// login function
// @Summary Login to the application
// @Description Login with email and password to get a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body auth.LoginRequest true "User information"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /auth/login [post]
func (h *authHandleApi) handleLogin(c *fiber.Ctx) error {
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

	res, err := h.client.LoginUser(c.Context(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.JSON(response.Response{
		Message:    "Success",
		StatusCode: fiber.StatusOK,
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
func (h *authHandleApi) handleRefreshToken(c *fiber.Ctx) error {
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

	res, err := h.client.RefreshToken(c.Context(), &pb.RefreshTokenRequest{
		RefreshToken: body.RefreshToken,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.JSON(response.Response{
		Message:    "Success",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
