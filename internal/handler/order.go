package handler

import (
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initOrderGroup(api *fiber.App) {
	order := api.Group("/api/order")

	order.Get("/hello", h.handlerHelloOrder)

	order.Use(middleware.Protector())

	order.Get("/", h.handleOrderAll)
	order.Post("/create", h.handlerOrderCreate)
	order.Get("/orderbyuser", h.handleOrderByUserId)
	order.Get("/:id", h.handleOrderById)

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
	authorization := c.Get("Authorization")

	us := authorization[7:]

	id, err := h.tokenManager.ValidateToken(us)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
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

	authorization := c.Get("Authorization")

	us := authorization[7:]

	id, err := h.tokenManager.ValidateToken(us)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
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
