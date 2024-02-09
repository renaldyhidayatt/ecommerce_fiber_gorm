package order

type OrderResponse struct {
	ID              uint                     `json:"id"`
	UserID          uint                     `json:"user_id"`
	Name            string                   `json:"name"`
	Phone           string                   `json:"phone"`
	Email           string                   `json:"email"`
	Courier         string                   `json:"courier"`
	ShippingMethod  string                   `json:"shipping_method"`
	ShippingCost    int                      `json:"shipping_cost"`
	TotalProduct    string                   `json:"total_product"`
	TotalPrice      int                      `json:"total_price"`
	TransactionID   string                   `json:"transaction_id"`
	OrderItems      []OrderItemResponse      `json:"order_items"`
	ShippingAddress *ShippingAddressResponse `json:"shipping_address"`
}

type OrderResponses struct {
	ID             uint   `json:"id"`
	UserID         uint   `json:"user_id"`
	UserName       string `json:"user_name"`
	UserEmail      string `json:"user_email"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Courier        string `json:"courier"`
	ShippingMethod string `json:"shipping_method"`
	ShippingCost   int    `json:"shipping_cost"`
	TotalProduct   string `json:"total_product"`
	TotalPrice     int    `json:"total_price"`
	TransactionID  string `json:"transaction_id"`
}

type OrderItemResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type ShippingAddressResponse struct {
	ID       uint   `json:"id"`
	Alamat   string `json:"alamat"`
	Provinsi string `json:"provinsi"`
	Negara   string `json:"negara"`
	Kota     string `json:"kota"`
}
