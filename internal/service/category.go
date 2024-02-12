package service

import (
	"ecommerce_fiber/internal/domain/requests/category"
	catresponse "ecommerce_fiber/internal/domain/response/category"
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/repository"
)

type categoryService struct {
	mapper     mapper.CategoryMapping
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository, mapper mapper.CategoryMapping) *categoryService {
	return &categoryService{repository: repository, mapper: mapper}
}

func (s *categoryService) GetCategories() ([]*catresponse.CategoryResponse, error) {
	res, err := s.repository.GetCategories()

	if err != nil {
		return nil, err
	}
	mapper := s.mapper.ToCategoryResponses(res)

	return mapper, nil
}

func (s *categoryService) GetCategory(categoryID int) (*catresponse.CategoryResponse, error) {

	res, err := s.repository.GetCategory(categoryID)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) GetCategorySlug(slug string) (*catresponse.CategoryResponse, error) {
	res, err := s.repository.GetCategorySlug(slug)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) CreateCategory(request *category.CreateCategoryRequest) (*catresponse.CategoryResponse, error) {

	res, err := s.repository.CreateCategory(request)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil

}

func (s *categoryService) UpdateCategory(updateCategory *category.UpdateCategoryRequest) (*catresponse.CategoryResponse, error) {

	category := category.UpdateCategoryRequest{
		ID:       updateCategory.ID,
		Name:     updateCategory.Name,
		FilePath: updateCategory.FilePath,
	}

	res, err := s.repository.UpdateCategory(&category)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) DeleteCategory(categoryID int) (*catresponse.CategoryResponse, error) {
	res, err := s.repository.DeleteCategory(categoryID)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil

}
