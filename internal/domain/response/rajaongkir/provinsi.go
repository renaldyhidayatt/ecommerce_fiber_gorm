package rajaongkir

type QueryProvinsi struct {
	ID string `json:"id"`
}

type StatusProvinsi struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type ResultsProvinsi struct {
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
}

type RajaOngkirResponseProvinsi struct {
	Rajaongkir struct {
		Query   []QueryProvinsi   `json:"query"`
		Status  StatusProvinsi    `json:"status"`
		Results []ResultsProvinsi `json:"results"`
	} `json:"rajaongkir"`
}
