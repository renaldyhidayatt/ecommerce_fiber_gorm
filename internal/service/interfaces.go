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
	rajaongkirresponse "ecommerce_fiber/internal/domain/response/rajaongkir"
	reviewres "ecommerce_fiber/internal/domain/response/review"
	userResponse "ecommerce_fiber/internal/domain/response/user"
	"ecommerce_fiber/internal/models"

	"github.com/midtrans/midtrans-go/snap"
)

type DashboardService interface {
	Dashboard() (*dashboard.DashboardResponse, error)
}

type AuthService interface {
	Register(input *auth.RegisterRequest) (*userResponse.UserResponse, error)
	Login(input *auth.LoginRequest) (*response.Token, error)
	RefreshToken(req auth.RefreshTokenRequest) (*response.Token, error)
}

type UserService interface {
	GetUsers() (*[]userResponse.UserResponse, error)
	CreateUser(request *auth.RegisterRequest) (*userResponse.UserResponse, error)
	GetUser(id int) (*userResponse.UserResponse, error)
	UpdateUser(request *user.UpdateUserRequest) (*userResponse.UserResponse, error)
	DeleteUser(id int) (*userResponse.UserResponse, error)
}

type CartService interface {
	FindAllByUserID(userID int) (*[]models.Cart, error)
	Create(cartRequest *cart.CartCreateRequest) (*models.Cart, error)
	Delete(cartID int) (*models.Cart, error)
	DeleteMany(cartIDs cart.DeleteCartRequest) (int64, error)
}

type CategoryService interface {
	GetCategories() ([]*catresponse.CategoryResponse, error)
	GetCategory(categoryID int) (*catresponse.CategoryResponse, error)
	GetCategorySlug(slug string) (*catresponse.CategoryResponse, error)
	CreateCategory(request *category.CreateCategoryRequest) (*catresponse.CategoryResponse, error)
	UpdateCategory(updateCategory *category.UpdateCategoryRequest) (*catresponse.CategoryResponse, error)
	DeleteCategory(categoryID int) (*catresponse.CategoryResponse, error)
}

type MidtransService interface {
	CreateTransaction(request *midtransrequest.CreateMidtransRequest) (*snap.Response, error)
}

type ProductService interface {
	GetProducts() ([]*pro.ProductResponse, error)
	CreateProduct(request *product.CreateProductRequest) (*pro.ProductResponse, error)
	GetProduct(productID int) (*pro.ProductResponse, error)
	GetProductSlug(slug string) (*pro.ProductResponse, error)
	UpdateProduct(request *product.UpdateProductRequest) (*pro.ProductResponse, error)
	DeleteProduct(productID int) (*pro.ProductResponse, error)
}

type OrderService interface {
	GetAll() (*[]or.OrderResponses, error)
	GetByID(orderID int) (*or.OrderResponse, error)
	CreateOrder(request *order.CreateOrderRequest) (*models.Order, error)
	OrdersByUser(userID int) (*[]models.Order, error)
}

type RajaOngkirService interface {
	GetProvinsi() (*rajaongkirresponse.RajaOngkirResponseProvinsi, error)
	GetCity(idProv int) (*rajaongkirresponse.RajaOngkirCityResponse, error)
	GetCost(request rajaongkirrequest.OngkosRequest) (*rajaongkirresponse.RajaOngkirOngkosResponse, error)
}

type ReviewService interface {
	GetAllReviews() (*[]reviewres.ReviewResponse, error)
	GetReviewByID(reviewID int) (*reviewres.ReviewResponse, error)
	CreateReview(request *review.CreateReviewRequest) (*reviewres.ReviewResponse, error)
}

type SliderService interface {
	GetAllSliders() (*[]models.Slider, error)
	GetSliderByID(sliderID int) (*models.Slider, error)
	CreateSlider(request slider.CreateSliderRequest) (*models.Slider, error)
	UpdateSliderByID(request slider.UpdateSliderRequest) (*models.Slider, error)
	DeleteSliderByID(sliderID int) (*models.Slider, error)
}
