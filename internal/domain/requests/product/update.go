package product

import "github.com/go-playground/validator/v10"

type UpdateProductRequest struct {
	Name        string `json:"name"`
	CategoryID  string `json:"category_id"`
	Description string `json:"description"`
	Price       int    `json:"price"`

	CountInStock int    `json:"count_in_stock"`
	Weight       int    `json:"weight"`
	Rating       int    `json:"rating"`
	Brand        string `json:"brand"`
	FilePath     string `json:"file_path"`
}

func (l *UpdateProductRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
