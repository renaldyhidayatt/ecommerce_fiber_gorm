package mapper

import (
	"ecommerce_fiber/internal/domain/response/category"
	"ecommerce_fiber/internal/models"
)

type categoryMapper struct {
}

func NewCategoryMapper() *categoryMapper {
	return &categoryMapper{}
}

func (m *categoryMapper) ToCategoryResponse(request *models.Category) *category.CategoryResponse {
	return &category.CategoryResponse{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		Slug:        request.SlugCategory,
		ImagePath:   request.ImageCategory,
		CreatedAt:   request.CreatedAt.String(),
		UpdatedAt:   request.UpdatedAt.String(),
	}
}

func (m *categoryMapper) ToCategoryResponses(requests *[]models.Category) []*category.CategoryResponse {
	var responses []*category.CategoryResponse
	for _, request := range *requests {
		responses = append(responses, m.ToCategoryResponse(&request))
	}
	return responses
}
