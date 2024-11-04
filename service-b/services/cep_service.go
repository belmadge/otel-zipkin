package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Location struct {
	City string `json:"localidade"`
}

func GetLocationByCEP(cep string) (*Location, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid or not found zipcode")
	}

	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	if location.City == "" {
		return nil, errors.New("invalid location data")
	}

	return &location, nil
}
