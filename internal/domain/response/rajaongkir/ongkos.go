package rajaongkir

type Query struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type CityDetails struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type CostDetail struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type Cost struct {
	Service     string       `json:"service"`
	Description string       `json:"description"`
	CostDetails []CostDetail `json:"cost"`
}

type CourierResult struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Costs []Cost `json:"costs"`
}

type RajaOngkirOngkosResponse struct {
	RajaOngkir struct {
		Query              Query           `json:"query"`
		Status             Status          `json:"status"`
		OriginDetails      CityDetails     `json:"origin_details"`
		DestinationDetails CityDetails     `json:"destination_details"`
		Results            []CourierResult `json:"results"`
	} `json:"rajaongkir"`
}
