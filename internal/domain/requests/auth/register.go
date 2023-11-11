package auth

import (
	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *RegisterRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}
