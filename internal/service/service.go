package service

import (
	"ecommerce_fiber/internal/repository"
	"ecommerce_fiber/pkg/auth"
	"ecommerce_fiber/pkg/hashing"
	"ecommerce_fiber/pkg/logger"
	"ecommerce_fiber/pkg/rajaongkir"
)

type Service struct {
	Cart       cartService
	Category   categoryService
	User       userService
	Midtrans   midtransService
	Order      orderService
	Product    productService
	RajaOngkir rajaOngkirService
	Review     reviewService
	Slider     sliderService
}

type Deps struct {
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	Token      auth.TokenManager
	Logger     logger.Logger
}

func NewService(deps Deps) *Service {
	return &Service{
		User:       *NewUserService(deps.Repository.User, deps.Hashing, deps.Token, deps.Logger),
		Category:   *NewCategoryService(deps.Repository.Category),
		Product:    *NewProductService(deps.Repository.Product),
		Cart:       *NewCartService(deps.Repository.Cart),
		Order:      *NewOrderService(deps.Repository.Order),
		Review:     *NewReviewService(deps.Repository.Review),
		RajaOngkir: *NewRajaOngkirService(rajaongkir.NewRajaOngkirAPI()),
	}
}
