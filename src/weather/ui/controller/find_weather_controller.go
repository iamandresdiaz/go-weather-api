package controller

import (
	"go-weather-api/src/weather/application/find"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type FindWeatherController struct {
	wf find.WeatherFinderInterface
}

func NewFindWeatherController(wf find.WeatherFinderInterface) *FindWeatherController {
	return &FindWeatherController{
		wf: wf,
	}
}

func (fwc *FindWeatherController) FindByCity(c *gin.Context) {
	city := strings.ToLower(c.Param("city"))

	weather, err := fwc.wf.FindByCity(city)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, weather.Weather())
}
