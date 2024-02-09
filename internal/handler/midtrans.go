package handler

import (
	midtransrequest "ecommerce_fiber/internal/domain/requests/midtrans_request"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initMidtransGroup(api *fiber.App) {
	midtrans := api.Group("/api/midtrans")

	midtrans.Post("/create-transaction", h.createTransaction)
}

func (h *Handler) createTransaction(c *fiber.Ctx) error {
	var request midtransrequest.CreateMidtransRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	res, err := h.services.Midtrans.CreateTransaction(&request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Transaction created", "data": res})
}
