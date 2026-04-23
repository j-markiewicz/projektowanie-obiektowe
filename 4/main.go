package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file: " + err.Error())
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/weather", getWeather)
	e.GET("/weather/:latlon", getWeather)

	if err := e.Start(":8000"); err != nil {
		panic("failed to start server:\n" + err.Error())
	}
}

// GET /weather
// GET /weather/:latlon
func getWeather(c *echo.Context) error {
	latlon := c.ParamOr("latlon", "50.0294913,19.9062002")
	coords := strings.Split(latlon, ",")

	if len(coords) != 2 {
		return c.String(http.StatusNotFound, "coordinates in unexpected format")
	}

	if strings.ContainsFunc(coords[0], func(r rune) bool { return !strings.ContainsRune("0.123456789", r) }) {
		return c.String(http.StatusNotFound, "invalid character in coordinates")
	}

	if strings.ContainsFunc(coords[1], func(r rune) bool { return !strings.ContainsRune("0.123456789", r) }) {
		return c.String(http.StatusNotFound, "invalid character in coordinates")
	}

	weatherUrl := "https://api.open-meteo.com/v1/forecast?latitude=" + coords[0] + "&longitude=" + coords[1] + "&hourly=temperature_2m,relative_humidity_2m,precipitation_probability,cloud_cover,surface_pressure,wind_speed_10m,wind_direction_10m,wind_gusts_10m,sunshine_duration,is_day&timezone=auto&past_days=1&forecast_days=7"

	res, err := http.Get(weatherUrl)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	weather, err := io.ReadAll(res.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Add("Content-Type", "application/json")
	return c.String(http.StatusOK, string(weather))
}
