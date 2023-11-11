package service

import (
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/domain/requests/product"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"
	"errors"
	"fmt"
)

type productService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *productService {
	return &productService{
		repository: repository,
	}
}

func (s *productService) GetAllProduct() (*[]models.Product, error) {
	res, err := s.repository.GetAllProducts()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *productService) CreateProduct(request *product.CreateProductRequest) (*models.Product, error) {
	var schema product.CreateProductRequest

	schema.Name = request.Name
	schema.CategoryID = request.CategoryID
	schema.Description = request.Description
	schema.Price = request.Price
	schema.CountInStock = request.CountInStock
	schema.Weight = request.Weight
	schema.Rating = request.Rating
	schema.FilePath = request.FilePath

	res, err := s.repository.CreateProduct(&schema)

	if err != nil {
		return res, nil
	}

	return res, nil

}

func (s *productService) GetById(productID int) (*models.Product, error) {
	res, err := s.repository.GetProductByID(productID)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *productService) GetBySlug(slug string) (*models.Product, error) {
	res, err := s.repository.GetProductBySlug(slug)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *productService) UpdateProduct(id int, request *product.UpdateProductRequest) (*models.Product, error) {

	var schema product.UpdateProductRequest

	schema.Name = request.Name
	schema.CategoryID = request.CategoryID
	schema.Description = request.Description
	schema.Price = request.Price
	schema.CountInStock = request.CountInStock
	schema.Weight = request.Weight
	schema.Rating = request.Rating
	schema.FilePath = request.FilePath

	res, err := s.repository.UpdateProduct(id, request)
	if err != nil {

		return nil, err
	}

	return res, nil
}

func (s *productService) DeleteProduct(productID int) (*models.Product, error) {
	res, err := s.repository.DeleteProduct(productID)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return nil, err
	}

	return res, nil
}

func (s *productService) UpdateQuantity(cart []*cart.CartCreateRequest) (bool, error) {
	if len(cart) == 0 {
		return false, errors.New("no cart data received")
	}

	for _, item := range cart {
		productID := item.ProductID
		quantity := item.Quantity

		product, err := s.GetById(productID)

		if err != nil {
			return false, err
		}

		currentStock := product.CountInStock

		newStock := currentStock - quantity

		result, err := s.repository.MyUpdateQuantity(productID, newStock)

		if err != nil {
			return false, err
		}
		if !result {
			return false, nil
		}
	}

	return true, nil
}
