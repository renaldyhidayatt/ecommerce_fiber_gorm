package rajaongkirrequest

import "github.com/go-playground/validator/v10"

type OngkosRequest struct {
	Asal   string `json:"asal"`
	Tujuan string `json:"tujuan"`
	Berat  int    `json:"berat"`
	Kurir  string `json:"kurir"`
}

func (l *OngkosRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
