package handler

import (
	"ecommerce_fiber/internal/domain/requests/category"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initCategoryGroup(api *fiber.App) {
	category := api.Group("/api/category")

	category.Get("/hello", h.handlerHelloCategory)
	category.Get("/", h.handleCategoryAll)

	category.Use(middleware.Protector())
	category.Get("/:id", h.handleCategoryById)
	category.Post("/create", h.handlerCategoryCreate)
	category.Put("/update/:id", h.handleCategoryUpdate)
	category.Delete("/delete/:id", h.handleCategoryDelete)
}

func (h *Handler) handlerHelloCategory(c *fiber.Ctx) error {
	return c.SendString("Handler Category")
}

func (h *Handler) handleCategoryAll(c *fiber.Ctx) error {
	res, err := h.services.Category.GetAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category data already to use",
		"status":  true,
		"data":    res,
	})
}

func (h *Handler) handleCategoryById(c *fiber.Ctx) error {
	catIDStr := c.Params("id")
	catId, err := strconv.Atoi(catIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid category ID",
			"error":   true,
		})
	}

	res, err := h.services.Category.GetByID(catId)

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

func (h *Handler) handlerCategoryCreate(c *fiber.Ctx) error {
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

	createReq := &category.CreateCategoryRequest{
		Name:     name,
		FilePath: imageURL,
	}

	res, err := h.services.Category.Create(createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "create successfuly", "data": res, "statusCode": fiber.StatusOK})
}

func (h *Handler) handleCategoryUpdate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File upload failed"})
	}

	userIDStr := c.Params("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
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

	updateReq := &category.UpdateCategoryRequest{
		Name:     name,
		FilePath: imageURL,
	}

	res, err := h.services.Category.UpdateByID(userID, updateReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update successfuly", "data": res, "statusCode": fiber.StatusOK})
}

func (h *Handler) handleCategoryDelete(c *fiber.Ctx) error {

	catIDStr := c.Params("id")
	categorId, err := strconv.Atoi(catIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid category ID",
			"error":   true,
		})
	}

	res, err := h.services.Category.DeleteByID(categorId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "successfully delete category",
		"data":       res,
		"statusCode": fiber.StatusOK,
	})
}
