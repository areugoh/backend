package main

import (
	"context"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ipinfo/go/v2/ipinfo"
)

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Name      string  `json:"name"`
	IP        string  `json:"-"`
}

func locationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loc := getLocationByIP(c.ClientIP())

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "location", loc))
		c.Next()
	}
}

func getLocationByIP(ip string) Location {
	mock := os.Getenv("MOCK")

	if mock == "true" {
		return Location{
			Latitude:  35.4729335,
			Longitude: 139.6146366,
			Name:      "Yokohama",
			IP:        ip,
		}
	}

	client := ipinfo.NewClient(nil, nil, os.Getenv("INFOIP_TOKEN"))
	info, _ := client.GetIPInfo(net.ParseIP(ip))

	longLat := strings.Split(info.Location, ",")
	lat, _ := strconv.ParseFloat(longLat[0], 64)
	lon, _ := strconv.ParseFloat(longLat[1], 64)

	return Location{
		Longitude: lon,
		Latitude:  lat,
		Name:      info.City,
		IP:        ip,
	}
}
