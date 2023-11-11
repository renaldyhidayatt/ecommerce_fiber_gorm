package cart

import "github.com/go-playground/validator/v10"

type DeleteCartRequest struct {
	CartIds []int `json:"cart_ids"`
}

func (l *DeleteCartRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
