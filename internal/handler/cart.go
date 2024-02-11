package handler

import (
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/domain/response"
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

// handleGetUserCarts function
// @Description Get all carts associated with the authenticated user
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /cart [get]
func (h *Handler) handleGetUserCarts(c *fiber.Ctx) error {
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

	res, err := h.services.Cart.FindAllByUserID(userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully get carts",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create cart
// @Description Create a new cart for the authenticated user
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body cart.CartCreateRequest true "Cart Create Request Body"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /cart/create [post]
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
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var body cart.CartCreateRequest

	body.UserID = userId

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

	newCart, err := h.services.Cart.Create(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully create cart",
		StatusCode: fiber.StatusOK,
		Data:       newCart,
	})
}

// @Summary Delete cart
// @Description Delete a cart by its ID
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param cart_id path int true "Cart ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /cart/delete/{cart_id} [delete]
func (h *Handler) handleCartDelete(c *fiber.Ctx) error {
	cartID, err := strconv.Atoi(c.Params("cart_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	deletedCart, err := h.services.Cart.Delete(cartID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}
	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully delete cart",
		StatusCode: fiber.StatusOK,
		Data:       deletedCart,
	})
}

// @Summary Delete many carts
// @Description Delete multiple carts by their IDs
// @Tags Cart
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body cart.DeleteCartRequest true "Delete Cart Request Body"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /api/cart/delete-many [delete]
func (h *Handler) handleDeleteManyCarts(c *fiber.Ctx) error {
	var deleteRequest cart.DeleteCartRequest
	if err := c.BodyParser(&deleteRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	deletedCount, err := h.services.Cart.DeleteMany(deleteRequest)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully delete carts",
		StatusCode: fiber.StatusOK,
		Data:       deletedCount,
	})
}
