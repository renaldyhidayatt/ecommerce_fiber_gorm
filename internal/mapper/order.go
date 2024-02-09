package mapper

import (
	"ecommerce_fiber/internal/domain/response/order"
	"ecommerce_fiber/internal/models"
)

type orderMapper struct {
}

func NewOrderMapper() *orderMapper {
	return &orderMapper{}
}

func (m *orderMapper) ToOrderResponse(requests *models.Order) *order.OrderResponse {
	return &order.OrderResponse{
		ID:              requests.ID,
		UserID:          requests.UserID,
		Name:            requests.Name,
		Phone:           requests.Phone,
		Email:           requests.Email,
		Courier:         requests.Courier,
		ShippingMethod:  requests.ShippingMethod,
		ShippingCost:    requests.ShippingCost,
		TotalProduct:    requests.TotalProduct,
		TotalPrice:      requests.TotalPrice,
		TransactionID:   requests.TransactionID,
		OrderItems:      m.ToOrderItemsResponses(requests.OrderItems),
		ShippingAddress: m.ToShippingAddressResponse(requests.ShippingAddress),
	}
}

func (m *orderMapper) ToOrderResponses(request *[]models.Order) []order.OrderResponses {
	var responses []order.OrderResponses

	for _, request := range *request {
		response := order.OrderResponses{
			ID:             request.ID,
			UserID:         request.UserID,
			UserName:       request.User.Name,
			UserEmail:      request.User.Email,
			Name:           request.Name,
			Phone:          request.Phone,
			Email:          request.Email,
			Courier:        request.Courier,
			ShippingMethod: request.ShippingMethod,
			ShippingCost:   request.ShippingCost,
			TotalProduct:   request.TotalProduct,
			TotalPrice:     request.TotalPrice,
			TransactionID:  request.TransactionID,
		}
		responses = append(responses, response)
	}

	return responses
}

func (m *orderMapper) ToOrderItemsResponses(requests []models.OrderItems) []order.OrderItemResponse {
	var responses []order.OrderItemResponse

	for _, request := range requests {
		response := order.OrderItemResponse{
			ID:       request.ID,
			Name:     request.Name,
			Quantity: request.Quantity,
			Price:    request.Price,
		}
		responses = append(responses, response)
	}

	return responses
}

func (m *orderMapper) ToShippingAddressResponse(requests models.ShippingAddress) *order.ShippingAddressResponse {
	return &order.ShippingAddressResponse{
		ID:       requests.ID,
		Alamat:   requests.Alamat,
		Provinsi: requests.Provinsi,
		Negara:   requests.Negara,
		Kota:     requests.Kota,
	}
}
