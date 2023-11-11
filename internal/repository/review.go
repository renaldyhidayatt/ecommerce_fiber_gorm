package repository

import (
	"ecommerce_fiber/internal/domain/requests/review"
	"ecommerce_fiber/internal/models"
	"errors"

	"gorm.io/gorm"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *reviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) GetAll() (*[]models.Review, error) {
	var reviews []models.Review

	db := r.db.Model(&reviews)

	checkReview := db.Debug().Find(&reviews)

	if checkReview.RowsAffected > 0 {
		return nil, errors.New("error")
	}

	return &reviews, nil
}

func (r *reviewRepository) GetByID(reviewID int) (*models.Review, error) {
	var review models.Review

	db := r.db.Model(&review)

	checkReview := db.Debug().Where("id = ?", reviewID)

	if checkReview.RowsAffected > 0 {
		return nil, errors.New("failed error")
	}

	return &review, nil

}

func (r *reviewRepository) CreateReview(request review.CreateReviewRequest, userID int, productID int) (*models.Review, error) {

	var product models.Product
	var review models.Review
	var user models.User

	dbProduct := r.db.Model(product)

	dbReview := r.db.Model(review)

	dbUser := r.db.Model(user)

	checkUser := dbUser.Debug().Where("id = ?", userID).First(&user)

	if checkUser.RowsAffected > 0 {
		return nil, errors.New("failed user id")
	}

	checkProduct := dbProduct.Debug().Where("id = ?", productID).First(&product)

	if checkProduct.RowsAffected > 0 {
		return nil, errors.New("failed product id")
	}

	checkReview := dbReview.Debug().Where("user_id = ? AND product_id", userID, productID).First(&review)

	if checkReview.RowsAffected > 0 {
		return nil, errors.New("failed review")
	}

	newReview := models.Review{
		Name:      checkUser.Name(),
		UserID:    uint(userID),
		Rating:    request.Rating,
		Comment:   request.Comment,
		ProductID: uint(productID),
	}

	if err := dbReview.Debug().Create(&newReview).Error; err != nil {
		return nil, err
	}

	var reviews []models.Review

	if err := r.db.Where("product_id = ?", productID).Find(&reviews).Error; err != nil {
		return nil, err
	}

	totalRating := 0
	for _, r := range reviews {
		totalRating += r.Rating
	}

	averageRating := float64(totalRating) / float64(len(reviews))

	if err := r.db.Model(&product).Update("rating", averageRating).Error; err != nil {
		return nil, err
	}

	return &review, nil
}
