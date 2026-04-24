package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

var ListWeather = make([]WeatherModel, 0)

func InitListWeather() {
	for lat := range 55 - 49 + 1 {
		for lon := range 24 - 14 + 1 {
			weatherUrl := "https://api.open-meteo.com/v1/forecast?latitude=" + strconv.Itoa(lat+49) + "&longitude=" + strconv.Itoa(lon+14) + "&hourly=temperature_2m,relative_humidity_2m,precipitation_probability,cloud_cover,surface_pressure,wind_speed_10m,wind_direction_10m,wind_gusts_10m,sunshine_duration,is_day&timezone=auto&past_days=1&forecast_days=7"

			res, err := http.Get(weatherUrl)
			if err != nil {
				continue
			}

			weatherStr, err := io.ReadAll(res.Body)
			if err != nil {
				continue
			}

			var weather Weather
			if err := json.Unmarshal([]byte(weatherStr), &weather); err != nil {
				continue
			}

			ListWeather = append(ListWeather, WeatherToModel(weather))
		}
	}
}

type Weather struct {
	Latitude             float32        `json:"latitude"`
	Longitude            float32        `json:"longitude"`
	GenerationtimeMs     float32        `json:"generationtime_ms"`
	UtcOffsetSeconds     int            `json:"utc_offset_seconds"`
	Timezone             string         `json:"timezone"`
	TimezoneAbbreviation string         `json:"timezone_abbreviation"`
	Elevation            float32        `json:"elevation"`
	HourlyUnits          HourlyUnits    `json:"hourly_units"`
	Hourly               HourlyForecast `json:"hourly"`
}

type WeatherModel struct {
	gorm.Model
	Latitude             float32
	Longitude            float32
	GenerationtimeMs     float32
	UtcOffsetSeconds     int
	Timezone             string
	TimezoneAbbreviation string
	Elevation            float32
	HourlyUnits          HourlyUnits           `gorm:"embedded;embeddedPrefix:hourly_units_"`
	Hourly               []HourlyForecastModel `gorm:"references:ForecastID"`
}

type HourlyUnits struct {
	Time                     string `json:"time"`
	Temperature2m            string `json:"temperature_2m"`
	RelativeHumidity2m       string `json:"relative_humidity_2m"`
	PrecipitationProbability string `json:"precipitation_probability"`
	CloudCover               string `json:"cloud_cover"`
	SurfacePressure          string `json:"surface_pressure"`
	WindSpeed10m             string `json:"wind_speed_10m"`
	WindDirection10m         string `json:"wind_direction_10m"`
	WindGusts10m             string `json:"wind_gusts_10m"`
	SunshineDuration         string `json:"sunshine_duration"`
	IsDay                    string `json:"is_day"`
}

type HourlyForecast struct {
	Time                     []string  `json:"time"`
	Temperature2m            []float32 `json:"temperature_2m"`
	RelativeHumidity2m       []int     `json:"relative_humidity_2m"`
	PrecipitationProbability []int     `json:"precipitation_probability"`
	CloudCover               []int     `json:"cloud_cover"`
	SurfacePressure          []float32 `json:"surface_pressure"`
	WindSpeed10m             []float32 `json:"wind_speed_10m"`
	WindDirection10m         []int     `json:"wind_direction_10m"`
	WindGusts10m             []float32 `json:"wind_gusts_10m"`
	SunshineDuration         []float32 `json:"sunshine_duration"`
	IsDay                    []int     `json:"is_day"`
}

type HourlyForecastModel struct {
	gorm.Model
	ForecastID               uint
	Time                     string
	Temperature2m            float32
	RelativeHumidity2m       int
	PrecipitationProbability int
	CloudCover               int
	SurfacePressure          float32
	WindSpeed10m             float32
	WindDirection10m         int
	WindGusts10m             float32
	SunshineDuration         float32
	IsDay                    int
}

func WeatherToModel(weather Weather) WeatherModel {
	model := WeatherModel{
		Latitude:             weather.Latitude,
		Longitude:            weather.Longitude,
		GenerationtimeMs:     weather.GenerationtimeMs,
		UtcOffsetSeconds:     weather.UtcOffsetSeconds,
		Timezone:             weather.Timezone,
		TimezoneAbbreviation: weather.TimezoneAbbreviation,
		Elevation:            weather.Elevation,
		HourlyUnits:          weather.HourlyUnits,
		Hourly:               make([]HourlyForecastModel, len(weather.Hourly.Time)),
	}

	for i := range len(weather.Hourly.Time) {
		model.Hourly = append(model.Hourly, HourlyForecastModel{
			Time:                     weather.Hourly.Time[i],
			Temperature2m:            weather.Hourly.Temperature2m[i],
			RelativeHumidity2m:       weather.Hourly.RelativeHumidity2m[i],
			PrecipitationProbability: weather.Hourly.PrecipitationProbability[i],
			CloudCover:               weather.Hourly.CloudCover[i],
			SurfacePressure:          weather.Hourly.SurfacePressure[i],
			WindSpeed10m:             weather.Hourly.WindSpeed10m[i],
			WindDirection10m:         weather.Hourly.WindDirection10m[i],
			WindGusts10m:             weather.Hourly.WindGusts10m[i],
			SunshineDuration:         weather.Hourly.SunshineDuration[i],
			IsDay:                    weather.Hourly.IsDay[i],
		})
	}

	return model
}

func WeatherFromModel(model WeatherModel) Weather {
	time := make([]string, len(model.Hourly))
	temperature2m := make([]float32, len(model.Hourly))
	relativeHumidity2m := make([]int, len(model.Hourly))
	precipitationProbability := make([]int, len(model.Hourly))
	cloudCover := make([]int, len(model.Hourly))
	surfacePressure := make([]float32, len(model.Hourly))
	windSpeed10m := make([]float32, len(model.Hourly))
	windDirection10m := make([]int, len(model.Hourly))
	windGusts10m := make([]float32, len(model.Hourly))
	sunshineDuration := make([]float32, len(model.Hourly))
	isDay := make([]int, len(model.Hourly))

	for i := range len(model.Hourly) {
		time = append(time, model.Hourly[i].Time)
		temperature2m = append(temperature2m, model.Hourly[i].Temperature2m)
		relativeHumidity2m = append(relativeHumidity2m, model.Hourly[i].RelativeHumidity2m)
		precipitationProbability = append(precipitationProbability, model.Hourly[i].PrecipitationProbability)
		cloudCover = append(cloudCover, model.Hourly[i].CloudCover)
		surfacePressure = append(surfacePressure, model.Hourly[i].SurfacePressure)
		windSpeed10m = append(windSpeed10m, model.Hourly[i].WindSpeed10m)
		windDirection10m = append(windDirection10m, model.Hourly[i].WindDirection10m)
		windGusts10m = append(windGusts10m, model.Hourly[i].WindGusts10m)
		sunshineDuration = append(sunshineDuration, model.Hourly[i].SunshineDuration)
		isDay = append(isDay, model.Hourly[i].IsDay)
	}

	return Weather{
		Latitude:             model.Latitude,
		Longitude:            model.Longitude,
		GenerationtimeMs:     model.GenerationtimeMs,
		UtcOffsetSeconds:     model.UtcOffsetSeconds,
		Timezone:             model.Timezone,
		TimezoneAbbreviation: model.TimezoneAbbreviation,
		Elevation:            model.Elevation,
		HourlyUnits:          model.HourlyUnits,
		Hourly: HourlyForecast{
			Time:                     time,
			Temperature2m:            temperature2m,
			RelativeHumidity2m:       relativeHumidity2m,
			PrecipitationProbability: precipitationProbability,
			CloudCover:               cloudCover,
			SurfacePressure:          surfacePressure,
			WindSpeed10m:             windSpeed10m,
			WindDirection10m:         windDirection10m,
			WindGusts10m:             windGusts10m,
			SunshineDuration:         sunshineDuration,
			IsDay:                    isDay,
		},
	}
}
