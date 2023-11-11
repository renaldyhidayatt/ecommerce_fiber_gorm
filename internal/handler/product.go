package handler

import (
	"ecommerce_fiber/internal/domain/requests/product"
	"ecommerce_fiber/internal/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initProductGroup(api *fiber.App) {
	product := api.Group("/api/product")

	product.Get("/hello", h.handlerProductHello)
	product.Get("/", h.handleProductAll)
	product.Get("/:id", h.handleProductById)

	product.Use(middleware.Protector())
	product.Post("/create", h.handleProductCreate)
	product.Put("/update/:id", h.handleProductUpdate)
	product.Delete("/delete/:id", h.handleProductDelete)

}

func (h *Handler) handlerProductHello(c *fiber.Ctx) error {
	return c.SendString("Handler Product")
}

func (h *Handler) handleProductAll(c *fiber.Ctx) error {
	res, err := h.services.Product.GetAllProduct()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.JSON(fiber.Map{
		"message": "product data already to use",
		"status":  true,
		"data":    res,
	})
}

func (h *Handler) handleProductById(c *fiber.Ctx) error {
	productIDStr := c.Params("id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
			"error":   true,
		})
	}

	res, err := h.services.Product.GetById(productId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "product data already to use",
		"statusCode": fiber.StatusOK,
		"data":       res,
	})

}

func (h *Handler) handleProductCreate(c *fiber.Ctx) error {
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid convert price",
			"error":   true,
		})
	}

	countInStock, err := strconv.Atoi(c.FormValue("count_in_stock"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid convert countInStock",
			"error":   true,
		})
	}

	weight, err := strconv.Atoi(c.FormValue("weight"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid convert weight",
			"error":   true,
		})
	}

	rating, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid convert rating",
			"error":   true,
		})
	}

	createReq := product.CreateProductRequest{
		Name:         c.FormValue("name"),
		CategoryID:   c.FormValue("category_id"),
		Description:  c.FormValue("description"),
		Price:        price,
		CountInStock: countInStock,
		Weight:       weight,
		Rating:       &rating,
	}

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

	createReq.FilePath = imageURL

	res, err := h.services.Product.CreateProduct(&createReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "create successfuly", "data": res, "statusCode": fiber.StatusOK})
}

func (h *Handler) handleProductUpdate(c *fiber.Ctx) error {

	productIDStr := c.Params("id")
	productId, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
			"error":   true,
		})
	}

	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid convert price",
			"error":   true,
		})
	}

	countInStock, err := strconv.Atoi(c.FormValue("count_in_stock"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid convert countInStock",
			"error":   true,
		})
	}

	weight, err := strconv.Atoi(c.FormValue("weight"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid convert weight",
			"error":   true,
		})
	}

	rating, err := strconv.Atoi(c.FormValue("rating"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid convert rating",
			"error":   true,
		})
	}

	updateReq := product.UpdateProductRequest{
		Name:         c.FormValue("name"),
		CategoryID:   c.FormValue("category_id"),
		Description:  c.FormValue("description"),
		Price:        price,
		CountInStock: countInStock,
		Weight:       weight,
		Rating:       rating,
	}

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

	updateReq.FilePath = imageURL

	res, err := h.services.Product.UpdateProduct(productId, &updateReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update successfuly", "data": res, "statusCode": fiber.StatusOK})
}

func (h *Handler) handleProductDelete(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
			"error":   true,
		})
	}

	res, err := h.services.Product.DeleteProduct(productId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":      true,
			"message":    err.Error(),
			"statusCode": fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "successfully delete product",
		"data":       res,
		"statusCode": fiber.StatusOK,
	})
}
