package mapper

import (
	"ecommerce_fiber/internal/domain/response/review"
	"ecommerce_fiber/internal/domain/response/user"
	"ecommerce_fiber/internal/models"
)

type reviewMapper struct{}

func NewReviewMapper() *reviewMapper {
	return &reviewMapper{}
}

func (m *reviewMapper) ToReviewResponse(request *models.Review) *review.ReviewResponse {
	return &review.ReviewResponse{
		ID:      int(request.ID),
		Name:    request.Name,
		Comment: request.Comment,
		Rating:  request.Rating,
		User: user.UserResponse{
			ID:         int(request.UserID),
			Name:       request.User.Name,
			Email:      request.User.Email,
			IsStaff:    request.User.IsStaff,
			Created_at: request.User.CreatedAt.String(),
			Updated_at: request.User.UpdatedAt.String(),
		},
		Sentiment: request.Sentiment,
		ProductID: int(request.ProductID),
	}
}

func (m *reviewMapper) ToReviewResponses(requests *[]models.Review) *[]review.ReviewResponse {
	var responses []review.ReviewResponse
	for _, request := range *requests {
		responses = append(responses, *m.ToReviewResponse(&request))
	}
	return &responses
}
