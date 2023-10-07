package main

import "github.com/gin-gonic/gin"

func getGraph(c *gin.Context) {
	long := c.Query("long")
	lat := c.Query("lat")

	c.JSON(200, gin.H{
		"long":    long,
		"lat":     lat,
		"message": "getGraph",
	})
}