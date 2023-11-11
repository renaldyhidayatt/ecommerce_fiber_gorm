package service

import (
	"ecommerce_fiber/internal/domain/requests/category"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"
)

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *categoryService {
	return &categoryService{repository: repository}
}

func (s *categoryService) GetAll() (*[]models.Category, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *categoryService) GetByID(categoryID int) (*models.Category, error) {

	res, err := s.repository.GetByID(categoryID)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *categoryService) GetBySlug(slug string) (*models.Category, error) {
	res, err := s.repository.GetBySlug(slug)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *categoryService) Create(request *category.CreateCategoryRequest) (*models.Category, error) {

	res, err := s.repository.Create(request)

	if err != nil {
		return nil, err
	}
	return res, nil

}

func (s *categoryService) UpdateByID(id int, updateCategory *category.UpdateCategoryRequest) (*models.Category, error) {

	category := category.UpdateCategoryRequest{

		Name:     updateCategory.Name,
		FilePath: updateCategory.FilePath,
	}

	res, err := s.repository.UpdateByID(id, &category)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *categoryService) DeleteByID(categoryID int) (*models.Category, error) {
	res, err := s.repository.DeleteByID(categoryID)

	if err != nil {
		return nil, err
	}

	return res, nil

}
