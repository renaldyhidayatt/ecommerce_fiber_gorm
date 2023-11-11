package handler

import (
	rajaongkirrequest "ecommerce_fiber/internal/domain/requests/rajaongkir_request"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initRajaongkirGroup(api *fiber.App) {
	rajaongkirGroup := api.Group("/api/rajaongkir")

	rajaongkirGroup.Get("/provinsi", h.handleRajaOngkirProvinsi)
	rajaongkirGroup.Get("/city/:id", h.handleRajaOngkirCity)
	rajaongkirGroup.Post("/cost", h.handleRajaOngkirCost)
}

func (h *Handler) handleRajaOngkirProvinsi(c *fiber.Ctx) error {
	provinsi, err := h.services.RajaOngkir.GetProvinsi()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(provinsi)
}

func (h *Handler) handleRajaOngkirCity(c *fiber.Ctx) error {
	id := c.Params("id")
	idProvinsi, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid province ID"})
	}
	city, err := h.services.RajaOngkir.GetCity(idProvinsi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(city)
}

func (h *Handler) handleRajaOngkirCost(c *fiber.Ctx) error {
	var request rajaongkirrequest.OngkosRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	cost, err := h.services.RajaOngkir.GetCost(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cost)
}
