package service

import (
	"ecommerce_fiber/internal/domain/response/dashboard"
	"ecommerce_fiber/internal/repository"
)

type dashboardService struct {
	user    repository.UserRepository
	product repository.ProductRepository
	order   repository.OrderRepository
}

func NewDashboardService(user repository.UserRepository, product repository.ProductRepository, order repository.OrderRepository) *dashboardService {
	return &dashboardService{
		user:    user,
		product: product,
		order:   order,
	}
}

func (s *dashboardService) Dashboard() (*dashboard.DashboardResponse, error) {
	totalUser, err := s.user.CountUser()

	if err != nil {
		return nil, err
	}

	totalProduct, err := s.product.CountProduct()

	if err != nil {
		return nil, err
	}

	totalOrder, err := s.order.CountOrder()

	if err != nil {
		return nil, err
	}

	pendapatan, err := s.order.SumTotalPrice()

	if err != nil {
		return nil, err
	}

	totalPendapatan, err := s.order.CalculateYearlyRevenue()

	if err != nil {
		return nil, err
	}

	return &dashboard.DashboardResponse{
		User:            totalUser,
		Product:         totalProduct,
		Order:           totalOrder,
		Pendapatan:      pendapatan,
		TotalPendapatan: totalPendapatan,
	}, nil
}
