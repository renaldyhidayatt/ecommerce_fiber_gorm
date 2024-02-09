package auth

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Confirm_password string `json:"confirm_password"`
}

func (l *RegisterRequest) Validate() error {
	if l.Password != l.Confirm_password {
		return errors.New("passwords do not match")
	}

	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}
