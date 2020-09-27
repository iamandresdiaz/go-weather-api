package persistence

import (
	"errors"
	"go-weather-api/src/weather/domain/entity"
	"go-weather-api/src/weather/domain/repository"
)

type WeatherRepository struct{}

var _ repository.WeatherRepository = &WeatherRepository{}

func NewWeatherRepository() *WeatherRepository {
	return &WeatherRepository{}
}

func (wr WeatherRepository) FindByCity(city string) (*entity.Weather, error) {
	if city != "barcelona" {
		return nil, errors.New("Weather not found")
	}

	weather := entity.Weather{
		ID:     "1baf7ca9-5063-4464-b3d4-25fb4d58f5c6",
		Value:  "17",
		Metric: "celcius",
		City:   "Barcelona",
	}

	return &weather, nil
}
