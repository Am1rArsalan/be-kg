package repo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Am1rArsalan/kelvin-green/response"
)

const openMeteoAPIURL = "https://api.open-meteo.com/v1/forecast"

type OpenMeteoRepoI interface {
	GetWeatherFromOpenMeteo(latitude, longitude string) (response.WeatherApiResponse, error)
}

type OpenMeteoRepo struct{}

func NewOpenMeteoRepo() *OpenMeteoRepo {
	return &OpenMeteoRepo{}
}

func (omr *OpenMeteoRepo) GetWeatherFromOpenMeteo(latitude, longitude string) (response.WeatherApiResponse, error) {
	apiURL := fmt.Sprintf("%s?latitude=%s&longitude=%s&hourly=temperature_2m,relativehumidity_2m&daily=weathercode,temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,sunrise,sunset&timezone=GMT", openMeteoAPIURL, latitude, longitude)

	resp, err := http.Get(apiURL)
	if err != nil {
		return response.WeatherApiResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response.WeatherApiResponse{}, fmt.Errorf("failed to fetch weather data, status code: %d", resp.StatusCode)
	}

	var weatherData response.WeatherApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return response.WeatherApiResponse{}, err
	}

	return weatherData, nil
}

