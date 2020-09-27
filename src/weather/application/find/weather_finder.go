package find

import (
	"go-weather-api/src/weather/domain/entity"
	"go-weather-api/src/weather/domain/repository"
)

type weatherFinder struct {
	wr repository.WeatherRepository
}

var _ WeatherFinderInterface = &weatherFinder{}

type WeatherFinderInterface interface {
	FindByCity(string) (*entity.Weather, error)
}

func (wf weatherFinder) FindByCity(city string) (*entity.Weather, error) {
	return wf.wr.FindByCity(city)
}
