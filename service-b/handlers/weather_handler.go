package handlers

import (
	"net/http"

	"github.com/belmadge/otel-zipkin/service-b/services"
	"github.com/belmadge/otel-zipkin/service-b/tracing"
	"github.com/gin-gonic/gin"
)

type WeatherResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func HandleWeather(c *gin.Context) {
	cep := c.Param("cep")

	ctx, span := tracing.StartSpan(c.Request.Context(), "cep_lookup")
	defer span.End()

	location, err := services.GetLocationByCEP(cep)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		return
	}

	ctx, span = tracing.StartSpan(ctx, "weather_lookup")
	defer span.End()

	weatherData, err := services.GetWeather(location.City)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "error retrieving weather"})
		return
	}

	tempF, tempK := services.ConvertTemperature(weatherData.TempC)

	response := WeatherResponse{
		City:  location.City,
		TempC: weatherData.TempC,
		TempF: tempF,
		TempK: tempK,
	}
	c.JSON(http.StatusOK, response)
}
