package handler

import (
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) initOrderGroup(api *fiber.App) {
	order := api.Group("/api/order")

	order.Get("/hello", h.handlerHelloOrder)

	order.Use(middleware.Protector())

	order.Get("/", h.handleOrderAll)
	order.Get("/create", h.handlerOrderCreate)
	order.Get("/:id", h.handleOrderById)
	order.Get("/orderbyuser", h.handleOrderByUserId)

}

func (h *Handler) handlerHelloOrder(c *fiber.Ctx) error {
	return c.SendString("Handler Order")
}

func (h *Handler) handleOrderAll(c *fiber.Ctx) error {
	res, err := h.services.Order.GetAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category data already to use",
		"status":  true,
		"data":    res,
	})
}

func (h *Handler) handleOrderById(c *fiber.Ctx) error {
	orderIdStr := c.Params("id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid order ID",
			"error":   true,
		})
	}

	res, err := h.services.Order.GetByID(orderId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "order data already to use",
		"statusCode": fiber.StatusOK,
		"data":       res,
	})
}

func (h *Handler) handleOrderByUserId(c *fiber.Ctx) error {
	userIdStr := c.Params("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid order ID",
			"error":   true,
		})
	}

	res, err := h.services.Order.OrdersByUser(userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "order data already to use",
		"statusCode": fiber.StatusOK,
		"data":       res,
	})
}

func (h *Handler) handlerOrderCreate(c *fiber.Ctx) error {

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid order ID",
			"error":   true,
		})
	}

	var body order.CreateOrderRequest

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

	res, err := h.services.Order.CreateOrder(userId, &body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "create successfuly", "data": res, "statusCode": fiber.StatusOK})
}
