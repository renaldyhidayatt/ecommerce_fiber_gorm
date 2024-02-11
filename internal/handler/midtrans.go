package handler

import (
	midtransrequest "ecommerce_fiber/internal/domain/requests/midtrans_request"
	"ecommerce_fiber/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initMidtransGroup(api *fiber.App) {
	midtrans := api.Group("/api/midtrans")

	midtrans.Post("/create-transaction", h.createTransaction)
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
func (h *Handler) createTransaction(c *fiber.Ctx) error {
	var request midtransrequest.CreateMidtransRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Midtrans.CreateTransaction(&request)
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
