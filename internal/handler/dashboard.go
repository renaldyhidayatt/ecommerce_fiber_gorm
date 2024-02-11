package handler

import (
	"ecommerce_fiber/internal/domain/response"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initDashboardGroup(api *fiber.App) {
	dashboard := api.Group("/api/dashboard")

	dashboard.Get("/", h.Dashboard)
}

// @Summary Get dashboard data
// @Description Get dashboard data
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /api/dashboard [get]
func (h *Handler) Dashboard(c *fiber.Ctx) error {
	dashboard, err := h.services.Dashboard.Dashboard()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}
	return c.JSON(response.Response{
		Message:    "dashboard data already to use",
		StatusCode: fiber.StatusOK,
		Data:       dashboard,
	})
}
