package api

import (
	midtransrequest "ecommerce_fiber/internal/domain/requests/midtrans_request"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type midtransHandleApi struct {
	client pb.MidtransServiceClient
}

func NewMidtransHandleApi(client pb.MidtransServiceClient, router *fiber.App) {
	midtransApi := &midtransHandleApi{
		client: client,
	}

	routerMidtrans := router.Group("/api/midtrans")

	routerMidtrans.Post("/create-transaction", midtransApi.CreateTransaction)
}

// @Summary Create transaction
// @Description Create a new transaction
// @Tags Midtrans
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createMidtransRequest body midtransrequest.CreateMidtransRequest true "Request body to create a new transaction"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /midtrans/create-transaction [post]
func (h *midtransHandleApi) CreateTransaction(c *fiber.Ctx) error {
	var body midtransrequest.CreateMidtransRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.CreateTransaction(c.Context(), &pb.CreateMidtransRequest{
		GrossAmount: int32(body.GrossAmount),
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       body.Email,
		Phone:       body.Phone,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "Transaction created successfully",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
