package handler

import (
	rajaongkirrequest "ecommerce_fiber/internal/domain/requests/rajaongkir_request"
	"ecommerce_fiber/internal/domain/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initRajaongkirGroup(api *fiber.App) {
	rajaongkirGroup := api.Group("/api/rajaongkir")

	rajaongkirGroup.Get("/provinsi", h.handleRajaOngkirProvinsi)
	rajaongkirGroup.Get("/city/:id", h.handleRajaOngkirCity)
	rajaongkirGroup.Post("/cost", h.handleRajaOngkirCost)
}

// @Summary Get provinces from RajaOngkir
// @Description Get list of provinces from RajaOngkir
// @Tags RajaOngkir
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorMessage
// @Router /rajaongkir/provinsi [get]
func (h *Handler) handleRajaOngkirProvinsi(c *fiber.Ctx) error {
	provinsi, err := h.services.RajaOngkir.GetProvinsi()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	return c.JSON(response.Response{
		Message:    "provinsi data already to use",
		StatusCode: fiber.StatusOK,
		Data:       provinsi,
	})
}

// @Summary Get cities by province ID from RajaOngkir
// @Description Get list of cities by province ID from RajaOngkir
// @Tags RajaOngkir
// @Param id path integer true "Province ID"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /rajaongkir/city/{id} [get]
func (h *Handler) handleRajaOngkirCity(c *fiber.Ctx) error {
	id := c.Params("id")
	idProvinsi, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}
	city, err := h.services.RajaOngkir.GetCity(idProvinsi)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	return c.JSON(response.Response{
		Message:    "city data already to use",
		StatusCode: fiber.StatusOK,
		Data:       city,
	})
}

// @Summary Get shipping cost from RajaOngkir
// @Description Get shipping cost from RajaOngkir
// @Tags RajaOngkir
// @Accept json
// @Produce json
// @Param request body rajaongkirrequest.OngkosRequest true "Request body"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /rajaongkir/cost [post]
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
