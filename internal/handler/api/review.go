package api

import (
	"ecommerce_fiber/internal/domain/requests/review"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/pb"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type reviewHandleApi struct {
	client pb.ReviewServiceClient
}

func NewReviewHandleApi(client pb.ReviewServiceClient, router *fiber.App) {
	reviewApi := &reviewHandleApi{
		client: client,
	}

	routerReview := router.Group("/api/review")

	routerReview.Get("/", reviewApi.GetReviews)
	routerReview.Get("/:id", reviewApi.GetReviewById)
	routerReview.Post("/create", reviewApi.CreateReview)
}

// @Summary Get all reviews
// @Description Get all reviews
// @Tags Review
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorMessage
// @Router /review [get]
func (h *reviewHandleApi) GetReviews(c *fiber.Ctx) error {
	res, err := h.client.GetReviews(c.Context(), &emptypb.Empty{})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "review data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
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
func (h *reviewHandleApi) GetReviewById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetReview(c.Context(), &pb.ReviewRequest{
		Id: int32(id),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "review data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
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
func (h *reviewHandleApi) CreateReview(c *fiber.Ctx) error {
	var request review.CreateReviewRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.CreateReview(c.Context(), &pb.CreateReviewRequest{
		ProductId: int32(request.ProductID),
		UserId:    int32(request.UserID),
		Rating:    int32(request.Rating),
		Comment:   request.Comment,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "review data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})

}
