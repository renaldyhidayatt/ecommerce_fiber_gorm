package category

import "github.com/go-playground/validator/v10"

type CreateCategoryRequest struct {
	Name     string `form:"name" json:"name"  validate:"required"`
	FilePath string `form:"file" json:"file" validate:"required"`
}

func (c *CreateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
