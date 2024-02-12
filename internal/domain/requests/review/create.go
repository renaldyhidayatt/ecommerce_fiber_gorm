package review

import "github.com/go-playground/validator/v10"

type CreateReviewRequest struct {
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}

func (l *CreateReviewRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
