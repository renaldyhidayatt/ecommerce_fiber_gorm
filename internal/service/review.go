package service

import (
	"ecommerce_fiber/internal/domain/requests/review"
	reviewres "ecommerce_fiber/internal/domain/response/review"
	"ecommerce_fiber/internal/mapper"
	"ecommerce_fiber/internal/repository"
)

type reviewService struct {
	reviewRepository repository.ReviewRepository
	mapper           mapper.ReviewMapping
}

func NewReviewService(reviewRepository repository.ReviewRepository, mapper mapper.ReviewMapping) *reviewService {
	return &reviewService{
		reviewRepository: reviewRepository,
		mapper:           mapper,
	}
}

func (s *reviewService) GetAllReviews() (*[]reviewres.ReviewResponse, error) {

	res, err := s.reviewRepository.GetAll()

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToReviewResponses(res)

	return mapper, nil
}

func (s *reviewService) GetReviewByID(reviewID int) (*reviewres.ReviewResponse, error) {

	res, err := s.reviewRepository.GetByID(reviewID)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToReviewResponse(res)

	return mapper, nil

}

func (s *reviewService) CreateReview(request *review.CreateReviewRequest) (*reviewres.ReviewResponse, error) {

	res, err := s.reviewRepository.CreateReview(*request)

	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToReviewResponse(res)

	return mapper, nil
}
