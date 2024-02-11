package handler

import (
	"ecommerce_fiber/internal/domain/requests/category"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initCategoryGroup(api *fiber.App) {
	category := api.Group("/api/category")

	category.Get("/hello", h.handlerHelloCategory)
	category.Get("/", h.handleCategoryAll)
	category.Get("/slug/:slug", h.handleCategorySlug)

	category.Use(middleware.Protector())
	category.Get("/:id", h.handleCategoryById)
	category.Post("/create", h.handlerCategoryCreate)
	category.Put("/update/:id", h.handleCategoryUpdate)
	category.Delete("/delete/:id", h.handleCategoryDelete)
}

// @Summary Get all categories
// @Description Get all categories
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /category [get]
func (h *Handler) handlerHelloCategory(c *fiber.Ctx) error {
	return c.SendString("Handler Category")
}

func (h *Handler) handleCategoryAll(c *fiber.Ctx) error {
	res, err := h.services.Category.GetAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "category data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get category by ID
// @Description Get a category by its ID
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /category/{id} [get]
func (h *Handler) handleCategoryById(c *fiber.Ctx) error {
	catIDStr := c.Params("id")
	catId, err := strconv.Atoi(catIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Category.GetByID(catId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "category data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get category by slug
// @Description Get a category by its slug
// @Tags Category
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param slug path string true "Category Slug"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /category/slug/{slug} [get]
func (h *Handler) handleCategorySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	res, err := h.services.Category.GetBySlug(slug)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "category data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create category
// @Description Create a new category
// @Tags Category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param name formData string true "Category Name"
// @Param image_category formData file true "Category Image"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /category/create [post]
func (h *Handler) handlerCategoryCreate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("image_category")
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
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.ErrorMessage{
				Error:      true,
				Message:    "Failed to upload file to Cloudinary " + err.Error(),
				StatusCode: fiber.StatusInternalServerError,
			},
		)
	}

	createReq := &category.CreateCategoryRequest{
		Name:     name,
		FilePath: imageURL,
	}

	res, err := h.services.Category.Create(createReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "category data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update category by ID
// @Description Update a category by its ID
// @Tags Category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param name formData string true "Category Name"
// @Param file formData file true "Category Image"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /category/update/{id} [put]
func (h *Handler) handleCategoryUpdate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    "File not found " + err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	userIDStr := c.Params("id")
	userID, err := strconv.Atoi(userIDStr)
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
			Message:    "Failed to upload file to Cloudinary " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	updateReq := &category.UpdateCategoryRequest{
		Name:     name,
		FilePath: imageURL,
	}

	res, err := h.services.Category.UpdateByID(userID, updateReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "category data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete category by ID
// @Description Delete a category by its ID
// @Tags Category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /category/delete/{id} [delete]
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
