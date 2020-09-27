package main

import (
	"go-weather-api/config/middleware"
	"go-weather-api/src/weather/infrastructure/persistence"
	"go-weather-api/src/weather/ui/controller"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	weatherRepository := persistence.NewWeatherRepository()
	findWeatherController := controller.NewFindWeatherController(weatherRepository)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Weather routes
	router.GET("/weather/:city", findWeatherController.FindByCity)

	//Starting the application
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080"
	}
	log.Fatal(router.Run(":" + appPort))
}
