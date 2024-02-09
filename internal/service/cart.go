package service

import (
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

type cartService struct {
	repository repository.CartRepository
	logger     logger.Logger
}

func NewCartService(repository repository.CartRepository, logger logger.Logger) *cartService {
	return &cartService{repository: repository, logger: logger}
}

func (s *cartService) FindAllByUserID(userID int) (*[]models.Cart, error) {
	res, err := s.repository.FindAllByUserID(userID)

	if err != nil {
		s.logger.Error("Error while getting cart by user id", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *cartService) Create(cartRequest *cart.CartCreateRequest) (*models.Cart, error) {

	fmt.Println("cart request", cartRequest.Weight)

	cart := &cart.CartCreateRequest{
		Name:         cartRequest.Name,
		Price:        cartRequest.Price,
		ImageProduct: cartRequest.ImageProduct,
		Quantity:     cartRequest.Quantity,
		Weight:       cartRequest.Weight,
		ProductID:    cartRequest.ProductID,
		UserID:       cartRequest.UserID,
	}

	res, err := s.repository.Create(cart)

	if err != nil {
		s.logger.Error("Error while creating cart", zap.Error(err))
		return nil, err
	}

	return res, err
}

func (s *cartService) Delete(cartID int) (*models.Cart, error) {

	res, err := s.repository.Delete(cartID)

	if err != nil {
		s.logger.Error("Error while deleting cart", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *cartService) DeleteMany(cartIDs cart.DeleteCartRequest) (int64, error) {
	res, err := s.repository.DeleteMany(cartIDs)

	if err != nil {
		s.logger.Error("Error while deleting cart many", zap.Error(err))
		return 0, err
	}

	return res, nil
}
