package mapper

import (
	"ecommerce_fiber/internal/domain/response/category"
	"ecommerce_fiber/internal/domain/response/order"
	"ecommerce_fiber/internal/domain/response/product"
	"ecommerce_fiber/internal/models"
)

type CategoryMapping interface {
	ToCategoryResponse(request *models.Category) *category.CategoryResponse
	ToCategoryResponses(requests *[]models.Category) []*category.CategoryResponse
}

type ProductMapping interface {
	ToProductResponse(request *models.Product) *product.ProductResponse
	ToProductResponses(requests *[]models.Product) []*product.ProductResponse
}

type OrderMapping interface {
	ToOrderResponse(requests *models.Order) *order.OrderResponse
	ToOrderResponses(request *[]models.Order) []order.OrderResponses
}
