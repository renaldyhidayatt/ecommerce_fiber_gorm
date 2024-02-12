package review

import "ecommerce_fiber/internal/domain/response/user"

type ReviewResponse struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Comment   string            `json:"comment"`
	Rating    int               `json:"rating"`
	User      user.UserResponse `json:"user"`
	Sentiment string            `json:"sentiment"`
	ProductID int               `json:"product_id"`
}
