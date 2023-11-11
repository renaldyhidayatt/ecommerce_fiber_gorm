package product

import "github.com/go-playground/validator/v10"

type CreateProductRequest struct {
	Name         string `json:"name"`
	CategoryID   string `json:"category_id"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Brand        string `json:"brand"`
	CountInStock int    `json:"count_in_stock"`
	Weight       int    `json:"weight"`
	Rating       *int   `json:"rating,omitempty"`
	FilePath     string `json:"file"`
}

func (l *CreateProductRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
