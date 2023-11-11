package service

import (
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"
)

type orderService struct {
	repository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *orderService {
	return &orderService{
		repository: orderRepository,
	}
}

func (s *orderService) GetAll() (*[]models.Order, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *orderService) CreateOrder(user_id int, request *order.CreateOrderRequest) (*models.Order, error) {
	res, err := s.repository.CreateOrder(user_id, request)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *orderService) GetByID(orderID int) (*models.Order, error) {

	res, err := s.repository.GetByID(orderID)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *orderService) OrdersByUser(userID int) (*[]models.Order, error) {
	res, err := s.repository.GetByUser(userID)

	if err != nil {
		return nil, err
	}

	return res, nil
}
