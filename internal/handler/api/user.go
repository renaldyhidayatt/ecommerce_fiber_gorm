package api

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/requests/user"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"ecommerce_fiber/internal/pb"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userHandleApi struct {
	client pb.UserServiceClient
}

func NewUserHandleApi(client pb.UserServiceClient, router *fiber.App) {
	api := &userHandleApi{
		client: client,
	}

	routerUser := router.Group("/api/user")

	routerUser.Get("/", api.handlerHello)

	routerUser.Use(middleware.Protector())

	routerUser.Get("/", api.handleUsers)
	routerUser.Get("/:id", api.handleUser)
	routerUser.Post("/create", api.handleUserCreate)
	routerUser.Put("/update/:id", api.handleUserUpdate)
	routerUser.Delete("/delete/:id", api.handleUserDelete)

}

func (h *userHandleApi) handlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler User")
}

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /user [get]
func (h *userHandleApi) handleUsers(c *fiber.Ctx) error {
	res, err := h.client.GetUsers(c.Context(), &emptypb.Empty{})

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
func (h *userHandleApi) handleUser(c *fiber.Ctx) error {
	userIDStr := c.Params("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetUser(c.Context(), &pb.UserRequest{Id: int64(userID)})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "successfully get user",
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
func (h *userHandleApi) handleUserCreate(c *fiber.Ctx) error {
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

	res, err := h.client.CreateUser(c.Context(), &pb.CreateUserRequest{
		Name:            body.Name,
		Email:           body.Email,
		Password:        body.Password,
		ConfirmPassword: body.Confirm_password,
	})

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
func (h *userHandleApi) handleUserUpdate(c *fiber.Ctx) error {
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

	res, err := h.client.UpdateUser(c.Context(), &pb.UpdateUserRequest{
		Id:       int64(userID),
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})

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
func (h *userHandleApi) handleUserDelete(c *fiber.Ctx) error {
	userIDStr := c.Params("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.DeleteUser(c.Context(), &pb.UserRequest{Id: int64(userID)})

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
