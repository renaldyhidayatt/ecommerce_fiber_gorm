package service

import (
	rajaongkirrequest "ecommerce_fiber/internal/domain/requests/rajaongkir_request"
	"ecommerce_fiber/pkg/rajaongkir"
	"encoding/json"
	"errors"
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

func (r *rajaOngkirService) GetProvinsi() (map[string]interface{}, error) {
	endpoint := "/starter/province"
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var result map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	return nil, errors.New("failed to fetch province data from RajaOngkir API. Status code: " + string(rune(res.StatusCode)))
}

func (r *rajaOngkirService) GetCity(idProv int) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/starter/city?province=%d", idProv)
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var result map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	return nil, errors.New("Failed to fetch city data from RajaOngkir API. Status code: " + string(rune(res.StatusCode)))
}

func (r *rajaOngkirService) GetCost(request rajaongkirrequest.OngkosRequest) (map[string]interface{}, error) {
	endpoint := "/starter/cost"
	client, headers := r.rajaOngkir.GetConnectionAndHeaders()
	url := fmt.Sprintf("%s%s", r.rajaOngkir.BaseURL, endpoint)

	payload := fmt.Sprintf("origin=%s&destination=%s&weight=%d&courier=%s",
		request.Asal, request.Tujuan, request.Berat, request.Kurir)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var result map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	return nil, errors.New("Failed to get shipping cost from RajaOngkir API. Status code: " + string(rune(res.StatusCode)))
}
