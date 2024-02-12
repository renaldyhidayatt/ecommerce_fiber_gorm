package service

import (
	"ecommerce_fiber/internal/domain/requests/order"
	or "ecommerce_fiber/internal/domain/response/order"
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/pkg/logger"

	"go.uber.org/zap"
)

type orderService struct {
	mapper     mapper.OrderMapping
	repository repository.OrderRepository
	logger     logger.Logger
}

func NewOrderService(orderRepository repository.OrderRepository, logger logger.Logger, mapper mapper.OrderMapping) *orderService {
	return &orderService{
		repository: orderRepository,
		logger:     logger,
		mapper:     mapper,
	}
}

func (s *orderService) GetAll() (*[]or.OrderResponses, error) {
	res, err := s.repository.GetAll()

	if err != nil {
		s.logger.Error("Error while getting all orders", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToOrderResponses(res)

	return &mapper, nil
}

func (s *orderService) CreateOrder(request *order.CreateOrderRequest) (*models.Order, error) {
	res, err := s.repository.CreateOrder(request)

	if err != nil {
		s.logger.Error("Error while creating order", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *orderService) GetByID(orderID int) (*or.OrderResponse, error) {

	res, err := s.repository.GetByID(orderID)

	if err != nil {
		s.logger.Error("Error while getting order by id", zap.Error(err))
		return nil, err
	}

	mapper := s.mapper.ToOrderResponse(res)

	return mapper, nil
}

func (s *orderService) OrdersByUser(userID int) (*[]models.Order, error) {
	res, err := s.repository.GetByUser(userID)

	if err != nil {
		s.logger.Error("Error while getting order by user id", zap.Error(err))
		return nil, err
	}

	return res, nil
}
