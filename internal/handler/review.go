package handler

import (
	"ecommerce_fiber/internal/domain/requests/review"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initReviewGroup(api *fiber.App) {
	reviewGroup := api.Group("/api/review")

	reviewGroup.Get("/", h.handleReviewsAll)
	reviewGroup.Get("/:id", h.handleReviewByID)

	reviewGroup.Use(middleware.Protector())
	reviewGroup.Post("/create", h.handleReviewCreate)
}

// @Summary Get all reviews
// @Description Get all reviews
// @Tags Review
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorMessage
// @Router /review [get]
func (h *Handler) handleReviewsAll(c *fiber.Ctx) error {
	reviews, err := h.services.Review.GetAllReviews()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	return c.JSON(response.Response{
		Message:    "review data already to use",
		StatusCode: fiber.StatusOK,
		Data:       reviews,
	})
}

// @Summary Get review by ID
// @Description Get review by ID
// @Tags Review
// @Param id path integer true "Review ID"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /review/{id} [get]
func (h *Handler) handleReviewByID(c *fiber.Ctx) error {
	reviewID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	review, err := h.services.Review.GetReviewByID(reviewID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	return c.JSON(response.Response{
		Message:    "review data already to use",
		StatusCode: fiber.StatusOK,
		Data:       review,
	})
}

// @Summary Create review
// @Description Create review
// @Tags Review
// @Param Authorization header string true "JWT token"
// @Param product_id path integer true "Product ID"
// @Accept json
// @Produce json
// @Param request body review.CreateReviewRequest true "Review data"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /review/create/{product_id} [post]
func (h *Handler) handleReviewCreate(c *fiber.Ctx) error {
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

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	var body review.CreateReviewRequest

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

	res, err := h.services.Review.CreateReview(productID, userId, &body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "review created successfully",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
