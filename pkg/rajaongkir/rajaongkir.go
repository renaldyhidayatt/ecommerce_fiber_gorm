package rajaongkir

import (
	"net/http"

	"github.com/spf13/viper"
)

type RajaOngkirAPI struct {
	BaseURL string
	ApiKey  string
}

func NewRajaOngkirAPI() *RajaOngkirAPI {
	return &RajaOngkirAPI{
		BaseURL: "api.rajaongkir.com",
		ApiKey:  viper.GetString("RAJAONGKIR_API"),
	}
}

func (api *RajaOngkirAPI) GetConnectionAndHeaders() (*http.Client, map[string]string) {
	client := &http.Client{}
	headers := map[string]string{
		"key":          api.ApiKey,
		"Content-Type": "application/x-www-form-urlencoded",
	}
	return client, headers
}
