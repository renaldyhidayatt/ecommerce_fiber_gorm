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

func (s *categoryService) GetAll() ([]*catresponse.CategoryResponse, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	mapper := s.mapper.ToCategoryResponses(res)

	return mapper, nil
}

func (s *categoryService) GetByID(categoryID int) (*catresponse.CategoryResponse, error) {

	res, err := s.repository.GetByID(categoryID)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) GetBySlug(slug string) (*catresponse.CategoryResponse, error) {
	res, err := s.repository.GetBySlug(slug)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) Create(request *category.CreateCategoryRequest) (*catresponse.CategoryResponse, error) {

	res, err := s.repository.Create(request)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil

}

func (s *categoryService) UpdateByID(id int, updateCategory *category.UpdateCategoryRequest) (*catresponse.CategoryResponse, error) {

	category := category.UpdateCategoryRequest{
		Name:     updateCategory.Name,
		FilePath: updateCategory.FilePath,
	}

	res, err := s.repository.UpdateByID(id, &category)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil
}

func (s *categoryService) DeleteByID(categoryID int) (*catresponse.CategoryResponse, error) {
	res, err := s.repository.DeleteByID(categoryID)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCategoryResponse(res)

	return mapper, nil

}
