package service

import (
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/domain/requests/product"
	pro "ecommerce_fiber/internal/domain/response/product"
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/pkg/logger"
	"errors"

	"go.uber.org/zap"
)

type productService struct {
	repository repository.ProductRepository
	logger     logger.Logger
	mapper     mapper.ProductMapping
}

func NewProductService(repository repository.ProductRepository, logger logger.Logger, mapper mapper.ProductMapping) *productService {
	return &productService{
		repository: repository,
		logger:     logger,
		mapper:     mapper,
	}
}

func (s *productService) GetAllProduct() ([]*pro.ProductResponse, error) {
	res, err := s.repository.GetAllProducts()

	if err != nil {
		s.logger.Error("Error while getting all products", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponses(res)

	return mapper, nil
}

func (s *productService) CreateProduct(request *product.CreateProductRequest) (*pro.ProductResponse, error) {
	var schema product.CreateProductRequest

	schema.Name = request.Name
	schema.CategoryID = request.CategoryID
	schema.Description = request.Description
	schema.Price = request.Price
	schema.CountInStock = request.CountInStock
	schema.Weight = request.Weight
	schema.Brand = request.Brand
	schema.Rating = request.Rating
	schema.FilePath = request.FilePath

	res, err := s.repository.CreateProduct(&schema)

	if err != nil {
		s.logger.Error("Error while creating product:", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil

}

func (s *productService) GetById(productID int) (*pro.ProductResponse, error) {
	res, err := s.repository.GetProductByID(productID)

	if err != nil {
		s.logger.Error("Error while getting product by id", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
}

func (s *productService) GetBySlug(slug string) (*pro.ProductResponse, error) {
	res, err := s.repository.GetProductBySlug(slug)

	if err != nil {
		s.logger.Error("Error while getting product by slug", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
}

func (s *productService) UpdateProduct(id int, request *product.UpdateProductRequest) (*pro.ProductResponse, error) {

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
		s.logger.Error("Error while updating product", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
}

func (s *productService) DeleteProduct(productID int) (*pro.ProductResponse, error) {
	res, err := s.repository.DeleteProduct(productID)
	if err != nil {
		s.logger.Error("Error while deleting product", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToProductResponse(res)

	return mapper, nil
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
			s.logger.Error("Error while getting product by id", zap.Error(err))
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
