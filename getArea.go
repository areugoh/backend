package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Condition string

const (
	ConditionGood         Condition = "good"
	ConditionModerate     Condition = "moderate"
	ConditionContaminated Condition = "contaminated"
	ConditionUnknown      Condition = "unknown"
	ConditionRadiation    Condition = "radiation"
	ConditionDanger       Condition = "danger"
	ConditionProtected    Condition = "protected"
)

type AquaticLocation struct {
	Name        string    `json:"name,omitempty"`
	Condition   Condition `json:"condition"`
	Temperature float64   `json:"temperature"`
	Distance    float64   `json:"distance"`
}

type AreaResponse struct {
	Location               Location        `json:"location"`
	NearestAquaticLocation AquaticLocation `json:"nearest_aquatic_location"`
}

func getArea(c *gin.Context) {
	location := c.Request.Context().Value("location").(Location)
	name, distance := getClosestSea(location)

	c.JSON(200, AreaResponse{
		Location: location,
		NearestAquaticLocation: AquaticLocation{
			Name:        name,
			Condition:   getAvgCondition(location),
			Temperature: getAvgTemperature(location),
			Distance:    distance,
		},
	})
}

func getClosestSea(location Location) (string, float64) {
	if os.Getenv("MOCK") == "true" {
		return "Tokyo Bay", 1.340 * 1.609
	}

	// API call to get {"distanceInMiles": 0} from KBGEO_URL with query parameters: lat=location.Latitude and lng=location.Longitude with header: "kb-auth-token": KBGEO_TOKEN
	req, _ := http.NewRequest("GET",
		fmt.Sprintf(
			"%s?lat=%v&lng=%v",
			os.Getenv("KBGEO_URL"),
			location.Latitude,
			location.Longitude,
		),
		nil)
	req.Header.Add("kb-auth-token", os.Getenv("KBGEO_TOKEN"))

	client := &http.Client{}
	resp, _ := client.Do(req)
	// {"distanceInMiles":1.347,"targetPoint":{"lat":35.47293350000000344834916177205741405487060546875,"lng":139.614636600000011412703315727412700653076171875},"coastlinePoint":{"lat":35.45830599999970189628584193997085094451904296875,"lng":139.6304440000000113286660052835941314697265625}}⏎

	type Response struct {
		DistanceInMiles float64 `json:"distanceInMiles"`
	}
	response := Response{}
	json.NewDecoder(resp.Body).Decode(&response)

	return "", response.DistanceInMiles * 1.609
}

func getAvgCondition(location Location) Condition {
	if os.Getenv("MOCK") == "true" {
		return ConditionModerate
	}

	//TODO: choose the condition with the highest count
	return ConditionUnknown
}

func getAvgTemperature(location Location) float64 {
	if os.Getenv("MOCK") == "true" {
		return 17.8
	}

	//TODO: API call to get the temperature
	return 0
}
