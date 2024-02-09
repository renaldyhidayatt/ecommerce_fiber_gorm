package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) initDashboardGroup(api *fiber.App) {
	dashboard := api.Group("/api/dashboard")

	dashboard.Get("/", h.Dashboard)
}

func (h *Handler) Dashboard(c *fiber.Ctx) error {
	dashboard, err := h.services.Dashboard.Dashboard()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    dashboard,
		"message": "success",
	})
}
