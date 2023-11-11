package auth

import (
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (l *LoginRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
