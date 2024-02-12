package mapper

import (
	"ecommerce_fiber/internal/domain/response/user"
	"ecommerce_fiber/internal/models"
)

type userMapping struct{}

func NewUserMapper() *userMapping {
	return &userMapping{}
}

func (m *userMapping) ToUserResponse(request *models.User) *user.UserResponse {
	return &user.UserResponse{
		ID:         int(request.ID),
		Name:       request.Name,
		Email:      request.Email,
		IsStaff:    request.IsStaff,
		Created_at: request.CreatedAt.String(),
		Updated_at: request.UpdatedAt.String(),
	}
}

func (m *userMapping) ToUserResponses(request *[]models.User) *[]user.UserResponse {
	var responses []user.UserResponse

	for _, request := range *request {
		response := user.UserResponse{
			ID:         int(request.ID),
			Name:       request.Name,
			Email:      request.Email,
			IsStaff:    request.IsStaff,
			Created_at: request.CreatedAt.String(),
			Updated_at: request.UpdatedAt.String(),
		}
		responses = append(responses, response)
	}

	return &responses
}
