package slider

import "github.com/go-playground/validator/v10"

type UpdateSliderRequest struct {
	Nama     string `json:"nama"`
	FilePath string `json:"file_path"`
}

func (l *UpdateSliderRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
