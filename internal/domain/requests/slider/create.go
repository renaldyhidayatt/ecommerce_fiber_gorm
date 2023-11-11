package slider

import "github.com/go-playground/validator/v10"

type CreateSliderRequest struct {
	Nama     string `json:"nama"`
	FilePath string `json:"file_path"`
}

func (l *CreateSliderRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
