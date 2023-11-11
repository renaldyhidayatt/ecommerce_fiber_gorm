package repository

import (
	"ecommerce_fiber/internal/domain/requests/order"
	"ecommerce_fiber/internal/models"
	"errors"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAll() (*[]models.Order, error) {
	var orders []models.Order

	db := r.db.Model(&orders)

	checkOrder := db.Debug().Find(&orders)

	if checkOrder.RowsAffected < 1 {
		return nil, errors.New("row kosong")
	}

	return &orders, nil
}

func (r *orderRepository) CreateOrder(user_id int, request *order.CreateOrderRequest) (*models.Order, error) {
	tx := r.db.Begin()

	var userModel models.User

	totalPrice, err := strconv.Atoi(request.TotalPrice)

	if err != nil {
		return nil, errors.New("error convert string to int")
	}

	dbUser := r.db.Model(userModel)

	checkUserByid := dbUser.Debug().Where("id = ?", user_id).First(&userModel)

	if checkUserByid.RowsAffected > 0 {
		return nil, errors.New("failed get user id")
	}

	orderCreate := models.Order{
		UserID:         userModel.ID,
		Name:           request.Name,
		Phone:          request.Phone,
		Courier:        request.Courier,
		Email:          userModel.Email,
		ShippingMethod: request.ShippingMethod,
		ShippingCost:   request.ShippingCost,
		TotalProduct:   request.TotalProduct,
		TotalPrice:     totalPrice,
		TransactionID:  uuid.New().String(),
	}

	if err := tx.Create(&orderCreate).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	shippingAddress := models.ShippingAddress{
		Alamat:   request.ShippingAddress.Alamat,
		Kota:     request.ShippingAddress.Kota,
		Negara:   request.ShippingAddress.Negara,
		Provinsi: request.ShippingAddress.Provinsi,
		OrderID:  orderCreate.ID,
	}

	if err := tx.Create(&shippingAddress).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range request.CartItems {
		orderItem := models.OrderItems{
			Name:     item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
			OrderID:  orderCreate.ID,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()
	return &orderCreate, nil
}

func (r *orderRepository) GetByID(orderID int) (*models.Order, error) {
	var order models.Order

	db := r.db.Model(&order)

	checkOrderbyId := db.Preload("OrderItems").Preload("ShippingAddress").Debug().Where("id = ?", orderID).First(&order)

	if checkOrderbyId.RowsAffected > 0 {
		return &order, errors.New("error not found order")
	}

	return &order, nil
}

func (r *orderRepository) GetByUser(userID int) (*[]models.Order, error) {
	var orders []models.Order

	db := r.db.Model(&orders)

	checkOrderbyUserId := db.Debug().Where("user_id = ?", userID).Find(&orders)

	if checkOrderbyUserId.RowsAffected > 0 {
		return &orders, errors.New("error user id")
	}

	return &orders, nil
}
