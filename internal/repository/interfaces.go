package repository

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/domain/requests/category"
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/domain/requests/product"
	"ecommerce_fiber/internal/domain/requests/review"
	"ecommerce_fiber/internal/domain/requests/slider"
	"ecommerce_fiber/internal/domain/requests/user"
	"ecommerce_fiber/internal/models"
)

type CartRepository interface {
	FindAllByUserID(userID int) (*[]models.Cart, error)
	Create(cartRequest *cart.CartCreateRequest) (*models.Cart, error)
	Delete(cartID int) (*models.Cart, error)
	DeleteMany(cartIDs cart.DeleteCartRequest) (int64, error)
}

type CategoryRepository interface {
	GetAll() (*[]models.Category, error)
	Create(request *category.CreateCategoryRequest) (*models.Category, error)
	GetByID(categoryID int) (*models.Category, error)
	GetBySlug(slug string) (*models.Category, error)
	UpdateByID(id int, updatedCategory *category.UpdateCategoryRequest) (*models.Category, error)
	DeleteByID(categoryID int) (*models.Category, error)
}

type OrderRepository interface {
	GetAll() (*[]models.Order, error)
	CreateOrder(user_id int, request *order.CreateOrderRequest) (*models.Order, error)
	GetByUser(userID int) (*[]models.Order, error)
	GetByID(orderID int) (*models.Order, error)
	CountOrder() (int, error)
	CalculateYearlyRevenue() ([]int, error)
	SumTotalPrice() (int, error)
}

type ProductRepository interface {
	GetAllProducts() (*[]models.Product, error)
	GetProductBySlug(slug string) (*models.Product, error)
	CreateProduct(request *product.CreateProductRequest) (*models.Product, error)
	GetProductByID(productID int) (*models.Product, error)
	MyUpdateQuantity(productID int, quantity int) (bool, error)
	UpdateProduct(productID int, request *product.UpdateProductRequest) (*models.Product, error)
	DeleteProduct(productID int) (*models.Product, error)
	CountProduct() (int, error)
}

type ReviewRepository interface {
	GetAll() (*[]models.Review, error)
	GetByID(reviewID int) (*models.Review, error)
	CreateReview(request review.CreateReviewRequest, userID int, productID int) (*models.Review, error)
}

type SliderRepository interface {
	GetAllSliders() (*[]models.Slider, error)
	GetSliderByID(sliderID int) (*models.Slider, error)
	CreateSlider(sliderRequest *slider.CreateSliderRequest) (*models.Slider, error)
	UpdateSliderByID(sliderID int, updatedSlider *slider.UpdateSliderRequest) (*models.Slider, error)
	DeleteSliderByID(sliderID int) (*models.Slider, error)
}
type UserRepository interface {
	CreateUser(request *auth.RegisterRequest) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserAll() (*[]models.User, error)
	GetUserById(id int) (*models.User, error)
	UpdateUserById(id int, updatedUser *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id int) (*models.User, error)
	CountUser() (int, error)
}
