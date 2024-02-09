package handler

import (
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initCartGroup(api *fiber.App) {
	cartGroup := api.Group("/api/cart")

	cartGroup.Use(middleware.Protector())

	cartGroup.Get("/", h.handleGetUserCarts)
	cartGroup.Post("/create", h.handleCartCreate)
	cartGroup.Delete("/delete/:cart_id", h.handleCartDelete)
	cartGroup.Delete("/delete-many", h.handleDeleteManyCarts)
}

func (h *Handler) handleGetUserCarts(c *fiber.Ctx) error {
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

	res, err := h.services.Cart.FindAllByUserID(userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "cart data already to use",
		"statusCode": fiber.StatusOK,
		"data":       res,
	})
}

func (h *Handler) handleCartCreate(c *fiber.Ctx) error {
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

	var body cart.CartCreateRequest

	body.UserID = userId

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

	newCart, err := h.services.Cart.Create(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":      false,
		"message":    "create successfully ",
		"data":       newCart,
		"statusCode": fiber.StatusOK,
	})
}

func (h *Handler) handleCartDelete(c *fiber.Ctx) error {
	cartID, err := strconv.Atoi(c.Params("cart_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
			"error":   true,
		})
	}

	deletedCart, err := h.services.Cart.Delete(cartID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "successfully delete product",
		"data":       deletedCart,
		"statusCode": fiber.StatusOK,
	})
}

func (h *Handler) handleDeleteManyCarts(c *fiber.Ctx) error {
	var deleteRequest cart.DeleteCartRequest
	if err := c.BodyParser(&deleteRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}

	deletedCount, err := h.services.Cart.DeleteMany(deleteRequest)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "successfully delete cart many",
		"data":       deletedCount,
		"statusCode": fiber.StatusOK,
	})
}
