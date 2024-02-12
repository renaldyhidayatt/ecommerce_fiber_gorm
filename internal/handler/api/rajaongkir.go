package api

import (
	rajaongkirrequest "ecommerce_fiber/internal/domain/requests/rajaongkir_request"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/pb"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type rajaOngkirHandleApi struct {
	client pb.RajaOngkirServiceClient
}

func NewRajaOngkirHandleApi(client pb.RajaOngkirServiceClient, router *fiber.App) {
	rajaApi := &rajaOngkirHandleApi{
		client: client,
	}

	rajaongkirGroup := router.Group("/api/rajaongkir")

	rajaongkirGroup.Get("/provinsi", rajaApi.handleProvinsi)
	rajaongkirGroup.Get("/city/:id", rajaApi.handleCity)
	rajaongkirGroup.Post("/cost", rajaApi.handleCost)

}

// @Summary Get provinces from RajaOngkir
// @Description Get list of provinces from RajaOngkir
// @Tags RajaOngkir
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorMessage
// @Router /rajaongkir/provinsi [get]
func (h *rajaOngkirHandleApi) handleProvinsi(c *fiber.Ctx) error {
	res, err := h.client.GetProvinsi(c.Context(), &emptypb.Empty{})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "province data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
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
func (h *rajaOngkirHandleApi) handleCity(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetCities(c.Context(), &pb.CityRequest{
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
		Message:    "city data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
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
func (h *rajaOngkirHandleApi) handleCost(c *fiber.Ctx) error {
	var request rajaongkirrequest.OngkosRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetCost(c.Context(), &pb.OngkosRequest{
		Asal:   request.Asal,
		Tujuan: request.Tujuan,
		Berat:  int32(request.Berat),
		Kurir:  request.Kurir,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "cost data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
