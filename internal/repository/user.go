package repository

import (
	"ecommerce_fiber/internal/domain/requests/auth"
	"ecommerce_fiber/internal/domain/requests/user"
	"ecommerce_fiber/internal/models"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(registerReq *auth.RegisterRequest) (*models.User, error) {
	var user models.User

	user.Name = registerReq.Name
	user.Email = registerReq.Email
	user.Password = registerReq.Password

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().Where("email = ?", registerReq.Email)

	if checkEmailExist.RowsAffected > 1 {

		return &user, errors.New("error")
	}

	result := db.Debug().Create(&user).Commit()

	if result.RowsAffected < 1 {
		return nil, errors.New("error create")
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email not found")
		}
		return nil, errors.New("error fetching email: " + err.Error())
	}

	return &user, nil
}

func (r *userRepository) GetUserAll() (*[]models.User, error) {
	var user []models.User

	db := r.db.Model(user)

	checkUSer := db.Debug().Find(&user)

	if checkUSer.RowsAffected < 1 {
		return nil, errors.New("failed to retrieve users: ")
	}

	return &user, nil
}

func (r *userRepository) GetUserById(id int) (*models.User, error) {

	var user models.User

	db := r.db.Model(user)

	checkCategoryById := db.Debug().Where("id = ?", id).First(&user)

	if checkCategoryById.RowsAffected > 0 {
		return &user, errors.New("error not found user")
	}

	return &user, nil
}

func (r *userRepository) UpdateUserById(id int, updatedUser *user.UpdateUserRequest) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	res, err := r.GetUserById(id)

	if err != nil {
		return nil, err
	}

	res.Name = updatedUser.Name
	res.Email = updatedUser.Email
	res.Password = updatedUser.Password

	if err := db.Save(&res).Error; err != nil {
		return nil, errors.New("error updating user: " + err.Error())
	}

	return res, nil
}

func (r *userRepository) DeleteUserById(id int) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	res, err := r.GetUserById(id)

	if err != nil {
		return nil, err
	}

	if err := db.Delete(&res).Error; err != nil {
		return nil, errors.New("error delete user: " + err.Error())
	}

	return res, nil

}
