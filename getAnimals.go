package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Animal struct {
	Name        string `json:"name"`
	InDanger    bool   `json:"in_danger"`
	Description string `json:"description"`
}

type AnimalsResponse struct {
	Location Location `json:"location"`
	Animals  []Animal `json:"animals"`
}

func getAnimals(c *gin.Context) {
	location := c.Request.Context().Value("location").(Location)

	c.JSON(200, AnimalsResponse{
		Location: location,
		Animals:  getAnimalList(location),
	})
}

func getAnimalList(location Location) []Animal {
	if os.Getenv("MOCK") == "true" {
		return []Animal{
			{
				Name:        "Dugong",
				InDanger:    true,
				Description: "The Dugong is a large marine mammal found in the waters near Yokohama. It is classified as endangered due to habitat destruction, pollution, and accidental entanglement in fishing nets. Human activities such as dredging, coastal development, and overfishing have led to the decline of their seagrass feeding grounds and disrupted their migration patterns.",
			},
			{
				Name:        "Japanese Devil Ray",
				InDanger:    true,
				Description: "The Japanese Devil Ray, also known as the Manta Ray, is facing a high risk of extinction. They are often caught as bycatch in fishing nets meant for other species. Additionally, pollution and habitat degradation from coastal development are threatening their populations. Conservation efforts are crucial to protect these majestic creatures in Yokohama's waters.",
			},
			{
				Name:        "Giant Pacific Octopus",
				InDanger:    false,
				Description: "The Giant Pacific Octopus is a fascinating species found in the waters near Yokohama. While it is not currently classified as endangered, it is important to monitor human activities to ensure their population remains stable. Overfishing and habitat destruction can still impact their ecosystem, so sustainable fishing practices and marine conservation measures play a vital role in maintaining their habitat and overall well-being.",
			},
		}
	}

	client := NewClient(os.Getenv("OPENAI_API_KEY"))

	req := CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: "you will be presented with a location or the name of a sea/ocean and your job is to provide a list of 3 aquatic animals living in that area with a flag of in_danger which should be true OR false meaning if it is a danger species and a short description of how humans affect to its daily life. Provide your answer in json format as follows: {name:'animal_name', in_danger:true, description:''}",
			},
			{
				Role:    "user",
				Content: location.Name,
			},
		},
		MaxTokens:   356,
		Temperature: 0,
	}

	resp, _ := client.CreateCompletions(req)
	fmt.Println(resp)

	animals := []Animal{}
	json.Unmarshal([]byte(resp.Choices[0].Message.Content), &animals)
	return animals
}
