package cart

import "github.com/go-playground/validator/v10"

type CartCreateRequest struct {
	Name         string `json:"name"`
	Price        string `json:"price"`
	ImageProduct string `json:"image_product"`
	Quantity     int    `json:"quantity"`
	ProductID    int    `json:"product_id"`
	UserID       int    `json:"user_id,omitempty"`
	Weight       int    `json:"weight"`
}

func (l *CartCreateRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
