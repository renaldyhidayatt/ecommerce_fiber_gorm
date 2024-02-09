package service

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/domain/requests/category"
	midtransrequest "ecommerce_fiber/internal/domain/requests/midtrans_request"
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/domain/requests/product"
	rajaongkirrequest "ecommerce_fiber/internal/domain/requests/rajaongkir_request"
	"ecommerce_fiber/internal/domain/requests/review"
	"ecommerce_fiber/internal/domain/requests/slider"
	"ecommerce_fiber/internal/domain/requests/user"
	"ecommerce_fiber/internal/domain/response"
	catresponse "ecommerce_fiber/internal/domain/response/category"
	"ecommerce_fiber/internal/domain/response/dashboard"
	or "ecommerce_fiber/internal/domain/response/order"
	pro "ecommerce_fiber/internal/domain/response/product"
	"ecommerce_fiber/internal/models"

	"github.com/midtrans/midtrans-go/snap"
)

type DashboardService interface {
	Dashboard() (*dashboard.DashboardResponse, error)
}

type AuthService interface {
	Register(input *auth.RegisterRequest) (*models.User, error)
	Login(input *auth.LoginRequest) (*response.Token, error)
	RefreshToken(req auth.RefreshTokenRequest) (*response.Token, error)
}

type UserService interface {
	GetUserAll() (*[]models.User, error)
	CreateUser(request *auth.RegisterRequest) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUserById(id int, request *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id int) (*models.User, error)
}

type CartService interface {
	FindAllByUserID(userID int) (*[]models.Cart, error)
	Create(cartRequest *cart.CartCreateRequest) (*models.Cart, error)
	Delete(cartID int) (*models.Cart, error)
	DeleteMany(cartIDs cart.DeleteCartRequest) (int64, error)
}

type CategoryService interface {
	GetAll() ([]*catresponse.CategoryResponse, error)
	GetByID(categoryID int) (*catresponse.CategoryResponse, error)
	GetBySlug(slug string) (*catresponse.CategoryResponse, error)
	Create(request *category.CreateCategoryRequest) (*catresponse.CategoryResponse, error)
	UpdateByID(id int, updateCategory *category.UpdateCategoryRequest) (*catresponse.CategoryResponse, error)
	DeleteByID(categoryID int) (*catresponse.CategoryResponse, error)
}

type MidtransService interface {
	CreateTransaction(request *midtransrequest.CreateMidtransRequest) (*snap.Response, error)
}

type ProductService interface {
	GetAllProduct() ([]*pro.ProductResponse, error)
	CreateProduct(request *product.CreateProductRequest) (*pro.ProductResponse, error)
	GetById(productID int) (*pro.ProductResponse, error)
	GetBySlug(slug string) (*pro.ProductResponse, error)
	UpdateProduct(id int, request *product.UpdateProductRequest) (*pro.ProductResponse, error)
	DeleteProduct(productID int) (*pro.ProductResponse, error)
}

type OrderService interface {
	GetAll() (*[]or.OrderResponses, error)
	GetByID(orderID int) (*or.OrderResponse, error)
	CreateOrder(user_id int, request *order.CreateOrderRequest) (*models.Order, error)
	OrdersByUser(userID int) (*[]models.Order, error)
}

type RajaOngkirService interface {
	GetProvinsi() (map[string]interface{}, error)
	GetCity(idProv int) (map[string]interface{}, error)
	GetCost(request rajaongkirrequest.OngkosRequest) (map[string]interface{}, error)
}

type ReviewService interface {
	GetAllReviews() (*[]models.Review, error)
	GetReviewByID(reviewID int) (*models.Review, error)
	CreateReview(productID int, user_id int, request *review.CreateReviewRequest) (*models.Review, error)
}

type SliderService interface {
	GetAllSliders() (*[]models.Slider, error)
	GetSliderByID(sliderID int) (*models.Slider, error)
	CreateSlider(request slider.CreateSliderRequest) (*models.Slider, error)
	UpdateSliderByID(sliderID int, request slider.UpdateSliderRequest) (*models.Slider, error)
	DeleteSliderByID(sliderID int) (*models.Slider, error)
}
