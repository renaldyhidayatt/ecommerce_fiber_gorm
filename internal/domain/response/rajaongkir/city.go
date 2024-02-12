package rajaongkir

type QueryCity struct {
	Province string `json:"province"`
	ID       string `json:"id"`
}

type StatusCity struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type ResultsCity struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type RajaOngkirCityResponse struct {
	RajaOngkir struct {
		Query   QueryCity     `json:"query"`
		Status  StatusCity    `json:"status"`
		Results []ResultsCity `json:"results"`
	} `json:"rajaongkir"`
}
