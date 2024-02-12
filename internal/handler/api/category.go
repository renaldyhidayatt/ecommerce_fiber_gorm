package api

import (
	"ecommerce_fiber/internal/domain/requests/category"
	"ecommerce_fiber/internal/domain/response"
	"ecommerce_fiber/internal/middleware"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/pkg/cloudinary"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryHandleApi struct {
	pb.UnimplementedCategoryServiceServer
	client     pb.CategoryServiceClient
	cloudinary cloudinary.MyCloudinary
}

func NewCategoryHandleApi(client pb.CategoryServiceClient, cloudinary cloudinary.MyCloudinary, router *fiber.App) {
	api := &categoryHandleApi{
		client:     client,
		cloudinary: cloudinary,
	}

	routerCategory := router.Group("/api/category")

	routerCategory.Get("/hello", api.handlerHello)
	routerCategory.Get("/", api.handleCategories)
	routerCategory.Get("/:id", api.handleCategory)
	routerCategory.Get("/slug/:slug", api.handleCategorySlug)

	routerCategory.Use(middleware.Protector())

	routerCategory.Post("/create", api.handleCategoryCreate)
	routerCategory.Put("/update/:id", api.handleCategoryUpdate)
	routerCategory.Delete("/delete/:id", api.handleCategoryDelete)

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
func (h *categoryHandleApi) handlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Category")
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
func (h *categoryHandleApi) handleCategories(c *fiber.Ctx) error {
	res, err := h.client.GetCategories(c.Context(), &emptypb.Empty{})

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
func (h *categoryHandleApi) handleCategory(c *fiber.Ctx) error {
	catIDStr := c.Params("id")
	catId, err := strconv.Atoi(catIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetCategory(c.Context(), &pb.CategoryRequest{
		Id: int64(catId),
	})

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
func (h *categoryHandleApi) handleCategorySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	res, err := h.client.GetCategorySlug(c.Context(), &pb.CategorySlugRequest{
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
func (h *categoryHandleApi) handleCategoryCreate(c *fiber.Ctx) error {
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

	res, err := h.client.CreateCategory(c.Context(), &pb.CreateCategoryRequest{
		Name:     name,
		FilePath: imageURL,
	})

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
func (h *categoryHandleApi) handleCategoryUpdate(c *fiber.Ctx) error {
	name := c.FormValue("name")

	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    "File not found " + err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	catIdStr := c.Params("id")

	catId, err := strconv.Atoi(catIdStr)

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

	res, err := h.client.UpdateCategory(c.Context(), &pb.UpdateCategoryRequest{
		Id:       int64(catId),
		Name:     updateReq.Name,
		FilePath: updateReq.FilePath,
	})

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
func (h *categoryHandleApi) handleCategoryDelete(c *fiber.Ctx) error {
	catIDStr := c.Params("id")

	categorId, err := strconv.Atoi(catIDStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.DeleteCategory(c.Context(), &pb.CategoryRequest{
		Id: int64(categorId),
	})

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
