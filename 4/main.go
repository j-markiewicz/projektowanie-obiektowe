package main

import (
	"os"

	ipgeolocation "github.com/IPGeolocation/ip-geolocation-go-sdk/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file: " + err.Error())
	}

	token := os.Getenv("IPGEOLOCATION_TOKEN")
	ipgeoClient, err = ipgeolocation.NewClient(&ipgeolocation.Config{
		APIKey: token,
	})

	if err != nil {
		panic("error creating ipgeolocation client" + err.Error())
	}

	InitListWeather()

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/api", GetApiWeather)
	e.GET("/api/", GetApiWeather)
	e.GET("/api/:latlon", GetApiWeather)
	e.POST("/api", GetApiWeather)
	e.POST("/api/", GetApiWeather)
	e.POST("/api/:latlon", GetApiWeather)

	e.GET("/list", GetListWeather)
	e.GET("/list/", GetListWeather)
	e.GET("/list/:latlon", GetListWeather)
	e.POST("/list", GetListWeather)
	e.POST("/list/", GetListWeather)
	e.POST("/list/:latlon", GetListWeather)

	if err := e.Start(":8000"); err != nil {
		panic("failed to start server:\n" + err.Error())
	}
}
