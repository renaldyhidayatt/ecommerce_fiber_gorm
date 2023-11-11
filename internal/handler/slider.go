package handler

import (
	"ecommerce_fiber/internal/domain/requests/slider"
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

func (h *Handler) handleSliderAll(c *fiber.Ctx) error {
	res, err := h.services.Slider.GetAllSliders()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.JSON(fiber.Map{
		"message": "slider data already to use",
		"status":  true,
		"data":    res,
	})
}

func (h *Handler) handleSliderById(c *fiber.Ctx) error {
	sliderIDStr := c.Params("id")
	sliderId, err := strconv.Atoi(sliderIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid slider ID",
			"error":   true,
		})
	}

	res, err := h.services.Slider.GetSliderByID(sliderId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "category data already to use",
		"statusCode": fiber.StatusOK,
		"data":       res,
	})
}

func (h *Handler) handlerSliderCreate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File upload failed"})
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open uploaded file"})
	}
	defer uploadedFile.Close()

	imageURL, err := h.cloudinary.UploadToCloudinary(uploadedFile, file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload file to Cloudinary " + err.Error()})
	}

	createReq := &slider.CreateSliderRequest{
		Nama:     name,
		FilePath: imageURL,
	}

	res, err := h.services.Slider.CreateSlider(*createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "create successfuly", "data": res, "statusCode": fiber.StatusOK})
}

func (h *Handler) handleSliderUpdate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File upload failed"})
	}

	sliderStr := c.Params("id")
	sliderId, err := strconv.Atoi(sliderStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid slider ID",
			"error":   true,
		})
	}

	uploadedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open uploaded file"})
	}
	defer uploadedFile.Close()

	imageURL, err := h.cloudinary.UploadToCloudinary(uploadedFile, file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upload file to Cloudinary " + err.Error()})
	}

	updateReq := &slider.UpdateSliderRequest{
		Nama:     name,
		FilePath: imageURL,
	}

	res, err := h.services.Slider.UpdateSliderByID(sliderId, *updateReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update successfuly", "data": res, "statusCode": fiber.StatusOK})
}

func (h *Handler) handleSliderDelete(c *fiber.Ctx) error {

	sliderIDStr := c.Params("id")
	sliderId, err := strconv.Atoi(sliderIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid slider ID",
			"error":   true,
		})
	}

	res, err := h.services.Slider.DeleteSliderByID(sliderId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "successfully delete slider",
		"data":       res,
		"statusCode": fiber.StatusOK,
	})
}
