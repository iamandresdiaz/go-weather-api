package repository

import (
	"go-weather-api/src/weather/domain/entity"
)

type WeatherRepository interface {
	FindByCity(string) (*entity.Weather, error)
}
