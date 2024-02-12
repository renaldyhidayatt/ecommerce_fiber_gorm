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

type sliderHandleApi struct {
	client     pb.SliderServiceClient
	cloudinary cloudinary.MyCloudinary
}

func NewSliderHandleApi(client pb.SliderServiceClient, cloudinary cloudinary.MyCloudinary, router *fiber.App) {
	sliderApi := &sliderHandleApi{
		client:     client,
		cloudinary: cloudinary,
	}

	sliderGroup := router.Group("/api/slider")

	sliderGroup.Get("/", sliderApi.GetSliders)

	sliderGroup.Use(middleware.Protector())

	sliderGroup.Get("/:id", sliderApi.GetSliderById)
	sliderGroup.Post("/create", sliderApi.CreateSlider)
	sliderGroup.Put("/update/:id", sliderApi.UpdateSlider)
	sliderGroup.Delete("/delete/:id", sliderApi.DeleteSlider)

}

// @Summary Get all sliders
// @Description Get all sliders
// @Tags Slider
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ErrorMessage
// @Failure 500 {object} response.ErrorMessage
// @Router /slider [get]
func (h *sliderHandleApi) GetSliders(c *fiber.Ctx) error {
	res, err := h.client.GetSliders(c.Context(), &emptypb.Empty{})

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
func (h *sliderHandleApi) GetSliderById(c *fiber.Ctx) error {
	sliderIDStr := c.Params("id")
	sliderId, err := strconv.Atoi(sliderIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.GetSlider(c.Context(), &pb.SliderRequest{
		Id: int64(sliderId),
	})

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
func (h *sliderHandleApi) CreateSlider(c *fiber.Ctx) error {
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

	createReq := &pb.CreateSliderRequest{
		Name:  name,
		Image: imageURL,
	}

	res, err := h.client.CreateSlider(c.Context(), createReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "slider created successfully",
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
func (h *sliderHandleApi) UpdateSlider(c *fiber.Ctx) error {
	sliderIDStr := c.Params("id")

	sliderId, err := strconv.Atoi(sliderIDStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

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

	updateReq := &pb.UpdateSliderRequest{
		Id:    int64(sliderId),
		Name:  name,
		Image: imageURL,
	}

	res, err := h.client.UpdateSlider(c.Context(), updateReq)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "slider updated successfully",
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
func (h *sliderHandleApi) DeleteSlider(c *fiber.Ctx) error {
	sliderIDStr := c.Params("id")

	sliderId, err := strconv.Atoi(sliderIDStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.client.DeleteSlider(c.Context(), &pb.SliderRequest{
		Id: int64(sliderId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response{
		Message:    "slider deleted successfully",
		StatusCode: fiber.StatusOK,
		Data:       res,
	})
}
