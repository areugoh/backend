package main

import "github.com/gin-gonic/gin"

type Condition string

const (
	ConditionGood Condition = "good"
	ConditionBad  Condition = "bad"
)

type AquaticLocation struct {
	Name        string    `json:"name"`
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

	c.JSON(200, AreaResponse{
		Location: location,
		NearestAquaticLocation: AquaticLocation{
			Name:        "Baltic Sea",
			Condition:   ConditionGood,
			Temperature: 17.8,
			Distance:    14.97,
		},
	})
}

func getClosestSea(location Location) (string, float64) {
	return "", 0
}

func getAvgCondition(location Location) float64 {
	return 0
}

func getAvgTemperature(location Location) float64 {
	return 0
}
