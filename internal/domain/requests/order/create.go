package order

import "github.com/go-playground/validator/v10"

type CartItemRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type ShippingAddressRequest struct {
	Alamat   string `json:"alamat"`
	Provinsi string `json:"provinsi"`
	Kota     string `json:"kota"`
	Negara   string `json:"negara"`
}

type CreateOrderRequest struct {
	Name            string                 `json:"name"`
	Phone           string                 `json:"phone"`
	Courier         string                 `json:"courier"`
	ShippingAddress ShippingAddressRequest `json:"shipping_address"`
	CartItems       []CartItemRequest      `json:"cart_items"`
	ShippingMethod  string                 `json:"shipping_method"`
	ShippingCost    int                    `json:"shipping_cost"`
	TotalProduct    string                 `json:"total_product"`
	TotalPrice      string                 `json:"total_price"`
}

func (l *CreateOrderRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}
