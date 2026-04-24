package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"
)

// GET /api[/[:latlon]]
func GetApiWeather(c *echo.Context) error {
	latlon := c.Param("latlon")

	if latlon == "" {
		latlon = GetUserLatLon(c.RealIP())
	}

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

	weatherStr, err := io.ReadAll(res.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var weather Weather
	if err := json.Unmarshal([]byte(weatherStr), &weather); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, weather)
}

// GET /list[/[:latlon]]
func GetListWeather(c *echo.Context) error {
	latlon := c.Param("latlon")

	if latlon == "" {
		latlon = GetUserLatLon(c.RealIP())
	}

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

	lat, err := strconv.ParseFloat(coords[0], 32)
	if err != nil {
		return c.String(http.StatusNotFound, "invalid coordinate")
	}

	lon, err := strconv.ParseFloat(coords[1], 32)
	if err != nil {
		return c.String(http.StatusNotFound, "invalid coordinate")
	}

	var closestWeather *WeatherModel = nil
	for _, weather := range ListWeather {
		if closestWeather == nil || Distance(float32(lat), float32(lon), weather.Latitude, weather.Longitude) < Distance(float32(lat), float32(lon), closestWeather.Latitude, closestWeather.Longitude) {
			closestWeather = &weather
		}
	}

	return c.JSON(http.StatusOK, WeatherFromModel(*closestWeather))
}
