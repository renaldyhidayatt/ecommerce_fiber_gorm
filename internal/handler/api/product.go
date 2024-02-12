package api

import (
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/pkg/cloudinary"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productHandeApi struct {
	client     pb.ProductServiceClient
	cloudinary cloudinary.MyCloudinary
}

func NewProductHandleApi(client pb.ProductServiceClient, cloudinary cloudinary.MyCloudinary, router *fiber.App) {
	productApi := &productHandeApi{
		client:     client,
		cloudinary: cloudinary,
	}

	routerProduct := router.Group("/api/product")

	routerProduct.Get("/hello", productApi.handlerHello)
	routerProduct.Get("/", productApi.handleGetProducts)
	routerProduct.Get("/slug/:slug", productApi.handleProductSlug)
	routerProduct.Get("/:id", productApi.handleGetProduct)

	routerProduct.Use(middleware.Protector())
	routerProduct.Post("/create", productApi.handleCreateProduct)
	routerProduct.Put("/update/:id", productApi.handleUpdateProduct)
	routerProduct.Delete("/delete/:id", productApi.handelDeleteProduct)

}

func (h *productHandeApi) handlerHello(c *fiber.Ctx) error {
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
func (h *productHandeApi) handleGetProducts(c *fiber.Ctx) error {
	res, err := h.client.GetProducts(c.Context(), &emptypb.Empty{})

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
func (h *productHandeApi) handleGetProduct(c *fiber.Ctx) error {
	productIDStr := c.Params("id")
	productId, err := strconv.Atoi(productIDStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetProduct(c.Context(), &pb.ProductRequest{
		Id: int64(productId),
	})

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
func (h *productHandeApi) handleProductSlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	res, err := h.client.GetProductSlug(c.Context(), &pb.ProductSlugRequest{
		Slug: slug,
	})

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
func (h *productHandeApi) handleCreateProduct(c *fiber.Ctx) error {
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

	createReq := pb.CreateProductRequest{
		Name:         c.FormValue("name"),
		CategoryId:   c.FormValue("category"),
		Description:  c.FormValue("description"),
		Brand:        c.FormValue("brand"),
		Price:        int32(price),
		CountInStock: int32(countInStock),
		Weight:       int32(weight),
		Rating:       int32(rating),
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

	res, err := h.client.CreateProduct(c.Context(), &createReq)

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
func (h *productHandeApi) handleUpdateProduct(c *fiber.Ctx) error {
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

	updateReq := pb.UpdateProductRequest{
		Id:           int32(productId),
		Name:         c.FormValue("name"),
		CategoryId:   c.FormValue("category"),
		Description:  c.FormValue("description"),
		Brand:        c.FormValue("brand"),
		Price:        int32(price),
		CountInStock: int32(countInStock),
		Weight:       int32(weight),
		Rating:       int32(rating),
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

	res, err := h.client.UpdateProduct(c.Context(), &updateReq)

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
func (h *productHandeApi) handelDeleteProduct(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.DeleteProduct(c.Context(), &pb.ProductRequest{
		Id: int64(productId),
	})

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
