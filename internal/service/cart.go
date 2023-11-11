package service

import (
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"
)

type cartService struct {
	repository repository.CartRepository
}

func NewCartService(repository repository.CartRepository) *cartService {
	return &cartService{repository: repository}
}

func (s *cartService) FindAllByUserID(userID int) (*[]models.Cart, error) {
	res, err := s.repository.FindAllByUserID(userID)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *cartService) Create(cartRequest *cart.CartCreateRequest) (*models.Cart, error) {

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
		return nil, err
	}

	return res, err
}

func (s *cartService) Delete(cartID int) (*models.Cart, error) {

	res, err := s.repository.Delete(cartID)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *cartService) DeleteMany(cartIDs cart.DeleteCartRequest) (int64, error) {
	res, err := s.repository.DeleteMany(cartIDs)

	if err != nil {
		return 0, err
	}

	return res, nil
}
