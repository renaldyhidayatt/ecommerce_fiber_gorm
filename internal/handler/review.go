package handler

import (
	"ecommerce_fiber/internal/domain/requests/review"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) initReviewGroup(api *fiber.App) {
	reviewGroup := api.Group("/api/review")

	reviewGroup.Get("/", h.handleReviewsAll)
	reviewGroup.Get("/:id", h.handleReviewByID)

	reviewGroup.Use(middleware.Protector())
	reviewGroup.Post("/create", h.handleReviewCreate)
}

func (h *Handler) handleReviewsAll(c *fiber.Ctx) error {
	reviews, err := h.services.Review.GetAllReviews()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(reviews)
}

func (h *Handler) handleReviewByID(c *fiber.Ctx) error {
	reviewID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid review ID"})
	}

	review, err := h.services.Review.GetReviewByID(reviewID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(review)
}

func (h *Handler) handleReviewCreate(c *fiber.Ctx) error {
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

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	var body review.CreateReviewRequest

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

	res, err := h.services.Review.CreateReview(productID, userId, &body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "create successfuly", "data": res, "statusCode": fiber.StatusOK})
}
