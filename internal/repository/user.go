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

func (r *userRepository) GetUsers() (*[]models.User, error) {
	var user []models.User

	db := r.db.Model(user)

	checkUSer := db.Debug().Find(&user)

	if checkUSer.RowsAffected < 1 {
		return nil, errors.New("failed to retrieve users: ")
	}

	return &user, nil
}

func (r *userRepository) GetUser(id int) (*models.User, error) {

	var user models.User

	db := r.db.Model(user)

	checkCategoryById := db.Debug().Where("id = ?", id).First(&user)

	if checkCategoryById.RowsAffected < 0 {
		return &user, errors.New("error not found user")
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(updatedUser *user.UpdateUserRequest) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	checkUser := db.Debug().Where("id = ?", updatedUser.ID).First(&user)

	if checkUser.RowsAffected < 1 {
		return nil, errors.New("error not found user")
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	updateUser := db.Debug().Updates(&user)

	if updateUser.RowsAffected < 1 {
		return nil, errors.New("error Failed update")
	}

	return &user, nil
}

func (r *userRepository) DeleteUser(id int) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	checkUser := db.Debug().Where("id = ?", id).First(&user)

	if checkUser.RowsAffected < 1 {
		return &user, errors.New("error not found user")
	}

	deleteUser := db.Debug().Delete(&user)

	if deleteUser.RowsAffected < 1 {
		return &user, errors.New("failed delete user")
	}

	return &user, nil

}

func (r *userRepository) CountUser() (int, error) {
	var user models.User

	db := r.db.Model(user)

	var totalUser int64

	db.Debug().Model(&user).Count(&totalUser)

	return int(totalUser), nil
}
