package handler

import (
	"ecommerce_fiber/internal/domain/requests/slider"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initSliderGroup(api *fiber.App) {
	slider := api.Group("/api/slider")

	slider.Get("/", h.handleSliderAll)

	slider.Use(middleware.Protector())

	slider.Get("/:id", h.handleSliderById)
	slider.Post("/create", h.handlerSliderCreate)
	slider.Put("/update/:id", h.handleSliderUpdate)
	slider.Delete("/delete/:id", h.handleSliderDelete)
}

// @Summary Get all sliders
// @Description Get all sliders
// @Tags Slider
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Failure 500 {object} response.ErrorMessage
// @Router /slider [get]
func (h *Handler) handleSliderAll(c *fiber.Ctx) error {
	res, err := h.services.Slider.GetAllSliders()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "slider data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get slider by ID
// @Description Get slider by ID
// @Tags Slider
// @Param id path integer true "Slider ID"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Failure 500 {object} response.ErrorMessage
// @Router /api/slider/{id} [get]
func (h *Handler) handleSliderById(c *fiber.Ctx) error {
	sliderIDStr := c.Params("id")
	sliderId, err := strconv.Atoi(sliderIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Slider.GetSliderByID(sliderId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "slider data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create a new slider
// @Description Create a new slider
// @Tags Slider
// @Accept multipart/form-data
// @Param name formData string true "Slider Name"
// @Param file formData file true "Slider Image File"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Failure 500 {object} response.ErrorMessage
// @Router /slider/create [post]
func (h *Handler) handlerSliderCreate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    "File not found " + err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    "Failed to open uploaded file " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	defer uploadedFile.Close()

	imageURL, err := h.cloudinary.UploadToCloudinary(uploadedFile, file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	createReq := &slider.CreateSliderRequest{
		Nama:     name,
		FilePath: imageURL,
	}

	res, err := h.services.Slider.CreateSlider(*createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully create slider",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update slider by ID
// @Description Update slider by ID
// @Tags Slider
// @Accept multipart/form-data
// @Param id path integer true "Slider ID"
// @Param name formData string true "Slider Name"
// @Param file formData file true "Slider Image File"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Failure 500 {object} response.ErrorMessage
// @Router /slider/update/{id} [put]
func (h *Handler) handleSliderUpdate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    "File not found " + err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	sliderStr := c.Params("id")
	sliderId, err := strconv.Atoi(sliderStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    "Failed to open uploaded file " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}
	defer uploadedFile.Close()

	imageURL, err := h.cloudinary.UploadToCloudinary(uploadedFile, file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	updateReq := &slider.UpdateSliderRequest{
		Nama:     name,
		FilePath: imageURL,
	}

	res, err := h.services.Slider.UpdateSliderByID(sliderId, *updateReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully update slider",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete slider by ID
// @Description Delete slider by ID
// @Tags Slider
// @Param id path integer true "Slider ID"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Failure 500 {object} response.ErrorMessage
// @Router /slider/delete/{id} [delete]
func (h *Handler) handleSliderDelete(c *fiber.Ctx) error {

	sliderIDStr := c.Params("id")
	sliderId, err := strconv.Atoi(sliderIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Slider.DeleteSliderByID(sliderId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "successfully delete slider",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
