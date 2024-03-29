package repository

import (
	"ecommerce_fiber/internal/domain/requests/cart"
	"ecommerce_fiber/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) FindAllByUserID(userID int) (*[]models.Cart, error) {

	var carts []models.Cart

	db := r.db.Model(&carts)

	checkCartByUserId := db.Debug().First(&carts, "user_id", userID)

	if checkCartByUserId.RowsAffected < 1 {
		return nil, errors.New("error not found cart")
	}

	return &carts, nil
}

func (r *cartRepository) Create(cartRequest *cart.CartCreateRequest) (*models.Cart, error) {
	var cartModel models.Cart

	fmt.Println("cart request", cartRequest.Weight)

	db := r.db.Model(&cartModel)

	cartModel.Name = cartRequest.Name
	cartModel.Price = cartRequest.Price
	cartModel.Image = cartRequest.ImageProduct
	cartModel.Quantity = cartRequest.Quantity
	cartModel.Weight = cartRequest.Weight
	cartModel.ProductID = uint(cartRequest.ProductID)
	cartModel.UserID = uint(cartRequest.UserID)

	addCart := db.Debug().Create(&cartModel).Commit()

	if addCart.RowsAffected < 1 {
		return nil, errors.New("error creating cart")
	}

	return &cartModel, nil
}

func (r *cartRepository) Delete(cartID int) (*models.Cart, error) {
	var cart models.Cart

	db := r.db.Model(cart)

	checkCartById := db.Debug().First(&cart, "id=?", cartID)

	if checkCartById.RowsAffected < 0 {
		return &cart, errors.New("error not found cart")
	}

	deleteCart := db.Debug().Delete(&cart)

	if deleteCart.Error != nil {
		return nil, deleteCart.Error
	}

	return &cart, nil
}

func (r *cartRepository) DeleteMany(cartIDs cart.DeleteCartRequest) (int64, error) {
	var cart models.Cart

	db := r.db.Model(cart)

	result := db.Debug().Where("id IN (?)", cartIDs).Delete(&cart)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil

}
