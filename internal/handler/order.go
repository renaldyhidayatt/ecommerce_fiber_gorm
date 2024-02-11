package handler

import (
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/domain/response"
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

// @Summary Greet the user for orders
// @Description Return a greeting message for orders
// @Tags Order
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /order/hello [get]
func (h *Handler) handlerHelloOrder(c *fiber.Ctx) error {
	return c.SendString("Handler Order")
}

// @Summary Get all orders
// @Description Retrieve all orders
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /order/ [get]
func (h *Handler) handleOrderAll(c *fiber.Ctx) error {
	res, err := h.services.Order.GetAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "order data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get order by ID
// @Description Retrieve an order by its ID
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /order/{id} [get]
func (h *Handler) handleOrderById(c *fiber.Ctx) error {
	orderIdStr := c.Params("id")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Order.GetByID(orderId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "order data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get orders by user ID
// @Description Retrieve orders associated with a specific user
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /order/by-user [get]
func (h *Handler) handleOrderByUserId(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	us := authorization[7:]

	id, err := h.tokenManager.ValidateToken(us)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Order.OrdersByUser(userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "order data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create order
// @Description Create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createOrderRequest body order.CreateOrderRequest true "Request body to create a new order"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /order/create [post]
func (h *Handler) handlerOrderCreate(c *fiber.Ctx) error {

	authorization := c.Get("Authorization")

	us := authorization[7:]

	id, err := h.tokenManager.ValidateToken(us)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var body order.CreateOrderRequest

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

	res, err := h.services.Order.CreateOrder(userId, &body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "order data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
