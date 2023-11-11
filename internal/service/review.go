package service

import (
	"ecommerce_fiber/internal/domain/requests/review"
	"ecommerce_fiber/internal/models"
	"ecommerce_fiber/internal/repository"
)

type reviewService struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewService(reviewRepository repository.ReviewRepository) *reviewService {
	return &reviewService{
		reviewRepository: reviewRepository,
	}
}

func (s *reviewService) GetAllReviews() (*[]models.Review, error) {

	res, err := s.reviewRepository.GetAll()

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *reviewService) GetReviewByID(reviewID int) (*models.Review, error) {

	res, err := s.reviewRepository.GetByID(reviewID)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *reviewService) CreateReview(productID int, user_id int, request *review.CreateReviewRequest) (*models.Review, error) {

	res, err := s.reviewRepository.CreateReview(*request, user_id, productID)

	if err != nil {
		return nil, err
	}

	return res, nil
}
