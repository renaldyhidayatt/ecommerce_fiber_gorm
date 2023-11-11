package auth

import "github.com/go-playground/validator/v10"

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (l *RefreshTokenRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
