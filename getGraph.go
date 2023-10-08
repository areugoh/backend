package main

import "github.com/gin-gonic/gin"

func getGraph(c *gin.Context) {
	location := c.Request.Context().Value("location").(Location)

	c.JSON(200, gin.H{
		"message":  "getGraph",
		"location": location,
	})
}
