package handler

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/requests/user"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initUserGroup(api *fiber.App) {
	user := api.Group("/api/user")

	user.Get("/", h.handlerUserAll)
	user.Get("/:id", h.handleUserById)

	user.Use(middleware.Protector())
	user.Post("/create", h.handleCreateUser)
	user.Put("/update/:id", h.handleUpdateUserById)
	user.Delete("/delete/:id", h.handleDeleteUserById)

}

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /user [get]
func (h *Handler) handlerUserAll(c *fiber.Ctx) error {
	res, err := h.services.User.GetUserAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "successfully get all user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags User
// @Param id path integer true "User ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /user/{id} [get]
func (h *Handler) handleUserById(c *fiber.Ctx) error {
	userIDStr := c.Params("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.User.GetUserById(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully get user by id",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body auth.RegisterRequest true "User details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /user/create [post]
func (h *Handler) handleCreateUser(c *fiber.Ctx) error {
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

	res, err := h.services.User.CreateUser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully create user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update user by ID
// @Description Update user by ID
// @Tags User
// @Param id path integer true "User ID"
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body user.UpdateUserRequest true "User details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /user/update/{id} [put]
func (h *Handler) handleUpdateUserById(c *fiber.Ctx) error {
	userIDStr := c.Params("id")
	var body user.UpdateUserRequest

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

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

	res, err := h.services.User.UpdateUserById(userID, &body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully update user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})

}

// @Summary Delete user by ID
// @Description Delete user by ID
// @Tags User
// @Param id path integer true "User ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /user/delete/{id} [delete]
func (h *Handler) handleDeleteUserById(c *fiber.Ctx) error {

	userIDStr := c.Params("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.User.DeleteUserById(userID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully delete user",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
