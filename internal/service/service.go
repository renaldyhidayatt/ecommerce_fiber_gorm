package service

import (
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/hashing"
	"ecommerce_fiber/pkg/logger"
	midtranspkg "ecommerce_fiber/pkg/midtrans_pkg"
	"ecommerce_fiber/pkg/rajaongkir"
)

type Service struct {
	Auth       AuthService
	Cart       CartService
	Category   CategoryService
	User       UserService
	Midtrans   MidtransService
	Order      OrderService
	Product    ProductService
	RajaOngkir RajaOngkirService
	Review     ReviewService
	Slider     SliderService
	Dashboard  DashboardService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	Token      auth.TokenManager
	Logger     logger.Logger
	Snap       *midtranspkg.SnapClient
	RajaOngkir *rajaongkir.RajaOngkirAPI
	Mapper     mapper.Mapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:       NewAuthService(deps.Repository.User, deps.Hashing, deps.Logger, deps.Token, deps.Mapper.UserMapper),
		User:       NewUserService(deps.Repository.User, deps.Hashing, deps.Logger, deps.Mapper.UserMapper),
		Category:   NewCategoryService(deps.Repository.Category, deps.Mapper.CategoryMapper),
		Product:    NewProductService(deps.Repository.Product, deps.Logger, deps.Mapper.ProductMapper),
		Cart:       NewCartService(deps.Repository.Cart, deps.Logger),
		Order:      NewOrderService(deps.Repository.Order, deps.Logger, deps.Mapper.OrderMapper),
		Midtrans:   NewMidtransService(deps.Snap),
		Review:     NewReviewService(deps.Repository.Review, deps.Mapper.ReviewMapper),
		Slider:     NewSliderService(deps.Repository.Slider, deps.Logger),
		RajaOngkir: NewRajaOngkirService(deps.RajaOngkir),
		Dashboard:  NewDashboardService(deps.Repository.User, deps.Repository.Product, deps.Repository.Order),
	}
}
