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
	GetCategories() (*[]models.Category, error)
	CreateCategory(request *category.CreateCategoryRequest) (*models.Category, error)
	GetCategory(categoryID int) (*models.Category, error)
	GetCategorySlug(slug string) (*models.Category, error)
	UpdateCategory(updatedCategory *category.UpdateCategoryRequest) (*models.Category, error)
	DeleteCategory(categoryID int) (*models.Category, error)
}

type OrderRepository interface {
	GetAll() (*[]models.Order, error)
	CreateOrder(request *order.CreateOrderRequest) (*models.Order, error)
	GetByUser(userID int) (*[]models.Order, error)
	GetByID(orderID int) (*models.Order, error)
	CountOrder() (int, error)
	CalculateYearlyRevenue() ([]int, error)
	SumTotalPrice() (int, error)
}

type ProductRepository interface {
	GetProducts() (*[]models.Product, error)
	GetProductSlug(slug string) (*models.Product, error)
	GetProduct(productID int) (*models.Product, error)
	CreateProduct(request *product.CreateProductRequest) (*models.Product, error)
	MyUpdateQuantity(productID int, quantity int) (bool, error)
	UpdateProduct(request *product.UpdateProductRequest) (*models.Product, error)
	DeleteProduct(productID int) (*models.Product, error)
	CountProduct() (int, error)
}

type ReviewRepository interface {
	GetAll() (*[]models.Review, error)
	GetByID(reviewID int) (*models.Review, error)
	CreateReview(request review.CreateReviewRequest) (*models.Review, error)
}

type SliderRepository interface {
	GetAllSliders() (*[]models.Slider, error)
	GetSliderByID(sliderID int) (*models.Slider, error)
	CreateSlider(sliderRequest *slider.CreateSliderRequest) (*models.Slider, error)
	UpdateSliderByID(updatedSlider *slider.UpdateSliderRequest) (*models.Slider, error)
	DeleteSliderByID(sliderID int) (*models.Slider, error)
}
type UserRepository interface {
	CreateUser(request *auth.RegisterRequest) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUsers() (*[]models.User, error)
	GetUser(id int) (*models.User, error)
	UpdateUser(updatedUser *user.UpdateUserRequest) (*models.User, error)
	DeleteUser(id int) (*models.User, error)
	CountUser() (int, error)
}
