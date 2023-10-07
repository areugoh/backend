package main

import "github.com/gin-gonic/gin"

func getAnimals(c *gin.Context) {
	long := c.Query("long")
	lat := c.Query("lat")

	c.JSON(200, gin.H{
		"long":    long,
		"lat":     lat,
		"message": "getAnimals",
	})
}