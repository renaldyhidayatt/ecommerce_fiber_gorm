package midtransrequest

import "github.com/go-playground/validator/v10"

type CreateMidtransRequest struct {
	GrossAmount int    `json:"gross_amount"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

func (l *CreateMidtransRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
