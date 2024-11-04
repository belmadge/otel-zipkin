package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type WeatherData struct {
	TempC float64 `json:"temp_c"`
}

func GetWeather(city string) (*WeatherData, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error retrieving weather data")
	}

	var data struct {
		Current WeatherData `json:"current"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data.Current, nil
}

func ConvertTemperature(tempC float64) (float64, float64) {
	tempF := tempC*1.8 + 32
	tempK := tempC + 273.15
	return tempF, tempK
}
