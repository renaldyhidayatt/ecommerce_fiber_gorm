package service

import (
	rajaongkirrequest "ecommerce_fiber/internal/domain/requests/rajaongkir_request"
	rajaongkirresponse "ecommerce_fiber/internal/domain/response/rajaongkir"
	"ecommerce_fiber/pkg/rajaongkir"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type rajaOngkirService struct {
	rajaOngkir *rajaongkir.RajaOngkirAPI
}

func NewRajaOngkirService(rajaOngkir *rajaongkir.RajaOngkirAPI) *rajaOngkirService {
	return &rajaOngkirService{rajaOngkir: rajaOngkir}
}

func (r *rajaOngkirService) GetProvinsi() (*rajaongkirresponse.RajaOngkirResponseProvinsi, error) {

	endpoint := "/starter/province"
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch province data from RajaOngkir API. Status code: %d", res.StatusCode)
	}

	var response rajaongkirresponse.RajaOngkirResponseProvinsi
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &response, nil
}

func (r *rajaOngkirService) GetCity(idProv int) (*rajaongkirresponse.RajaOngkirCityResponse, error) {
	var result rajaongkirresponse.RajaOngkirCityResponse

	endpoint := fmt.Sprintf("/starter/city?province=%d", idProv)
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch city data from RajaOngkir API. Status code: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &result, nil
}

func (r *rajaOngkirService) GetCost(request rajaongkirrequest.OngkosRequest) (*rajaongkirresponse.RajaOngkirOngkosResponse, error) {
	var result rajaongkirresponse.RajaOngkirOngkosResponse

	endpoint := "/starter/cost"
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	payload := fmt.Sprintf("origin=%s&destination=%s&weight=%d&courier=%s",
		request.Asal, request.Tujuan, request.Berat, request.Kurir)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get shipping cost from RajaOngkir API. Status code: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &result, nil
}
