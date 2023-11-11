package category

import "github.com/go-playground/validator/v10"

type UpdateCategoryRequest struct {
	Name     string `json:"name"`
	FilePath string `json:"file_path"`
}

func (l *UpdateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
