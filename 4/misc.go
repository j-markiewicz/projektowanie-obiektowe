package main

import (
	"context"
	"math"

	ipgeolocation "github.com/IPGeolocation/ip-geolocation-go-sdk/v2"
)

func Distance(lat1 float32, lon1 float32, lat2 float32, lon2 float32) float32 {
	return float32(math.Sqrt(float64((lat1-lat2)*(lat1-lat2) + (lon1-lon2)*(lon1-lon2))))
}

var ipgeoClient *ipgeolocation.Client

func GetUserLatLon(ip string) string {
	res, err := ipgeoClient.LookupIPGeolocation(context.Background(), &ipgeolocation.LookupRequest{
		IP: ip,
	})

	if err != nil {
		return "50.0294913,19.9062002"
	}

	info := res.Data

	if info.Location.Latitude == nil || info.Location.Longitude == nil {
		return "50.0294913,19.9062002"
	}

	return *info.Location.Latitude + "," + *info.Location.Longitude
}
