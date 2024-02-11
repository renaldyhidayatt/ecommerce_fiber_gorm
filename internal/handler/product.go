package handler

import (
	"ecommerce_fiber/internal/domain/requests/product"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initProductGroup(api *fiber.App) {
	product := api.Group("/api/product")

	product.Get("/hello", h.handlerProductHello)
	product.Get("/", h.handleProductAll)
	product.Get("/slug/:slug", h.handleProductFindBySlug)
	product.Get("/:id", h.handleProductById)

	product.Use(middleware.Protector())
	product.Post("/create", h.handleProductCreate)
	product.Put("/update/:id", h.handleProductUpdate)
	product.Delete("/delete/:id", h.handleProductDelete)

}

// @Summary Greet the user for products
// @Description Return a greeting message for products
// @Tags Product
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /product/hello [get]
func (h *Handler) handlerProductHello(c *fiber.Ctx) error {
	return c.SendString("Handler Product")
}

// @Summary Get all products
// @Description Retrieve all products
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /product [get]
func (h *Handler) handleProductAll(c *fiber.Ctx) error {
	res, err := h.services.Product.GetAllProduct()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.JSON(response.Response{
		Message:    "product data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Get product by ID
// @Description Retrieve a product by its ID
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /product/{id} [get]
func (h *Handler) handleProductById(c *fiber.Ctx) error {
	productIDStr := c.Params("id")
	productId, err := strconv.Atoi(productIDStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.GetById(productId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "product data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})

}

// @Summary Get product by slug
// @Description Retrieve a product by its slug
// @Tags Product
// @Produce json
// @Param slug path string true "Product Slug"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /product/slug/{slug} [get]
func (h *Handler) handleProductFindBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	res, err := h.services.Product.GetBySlug(slug)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "product data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Create product
// @Description Create a new product
// @Tags Product
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param name formData string true "Product Name"
// @Param category formData string true "Category ID"
// @Param description formData string true "Product Description"
// @Param brand formData string true "Product Brand"
// @Param price formData integer true "Product Price"
// @Param countInStock formData integer true "Product Count In Stock"
// @Param weight formData integer true "Product Weight"
// @Param rating formData integer true "Product Rating"
// @Param image_product formData file true "Product Image"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /product/create [post]
func (h *Handler) handleProductCreate(c *fiber.Ctx) error {
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	countInStock, err := strconv.Atoi(c.FormValue("countInStock"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	weight, err := strconv.Atoi(c.FormValue("weight"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	rating, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	createReq := product.CreateProductRequest{
		Name:         c.FormValue("name"),
		CategoryID:   c.FormValue("category"),
		Description:  c.FormValue("description"),
		Brand:        c.FormValue("brand"),
		Price:        price,
		CountInStock: countInStock,
		Weight:       weight,
		Rating:       &rating,
	}

	file, err := c.FormFile("image_product")
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
			Message:    err.Error(),
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

	createReq.FilePath = imageURL

	if err := createReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.CreateProduct(&createReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "product data already to use",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Update product
// @Description Update an existing product
// @Tags Product
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param name formData string true "Product Name"
// @Param category formData string true "Category ID"
// @Param description formData string true "Product Description"
// @Param brand formData string true "Product Brand"
// @Param price formData integer true "Product Price"
// @Param countInStock formData integer true "Product Count In Stock"
// @Param weight formData integer true "Product Weight"
// @Param rating formData integer true "Product Rating"
// @Param image_product formData file true "Product Image"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /product/update/{id} [put]
func (h *Handler) handleProductUpdate(c *fiber.Ctx) error {

	productIDStr := c.Params("id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	countInStock, err := strconv.Atoi(c.FormValue("countInStock"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	weight, err := strconv.Atoi(c.FormValue("weight"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	rating, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	updateReq := product.UpdateProductRequest{
		Name:         c.FormValue("name"),
		CategoryID:   c.FormValue("category"),
		Description:  c.FormValue("description"),
		Brand:        c.FormValue("brand"),
		CountInStock: countInStock,
		Price:        price,
		Weight:       weight,
		Rating:       rating,
	}

	file, err := c.FormFile("image_product")
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
			Message:    "Failed to upload file to Cloudinary " + err.Error(),
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	updateReq.FilePath = imageURL

	if err := updateReq.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.UpdateProduct(productId, &updateReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "update successfuly",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}

// @Summary Delete product
// @Description Delete an existing product
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Router /product/delete/{id} [delete]
func (h *Handler) handleProductDelete(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.DeleteProduct(productId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "delete successfuly",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
