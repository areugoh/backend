package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.Use(locationMiddleware())
	{
		v1.GET("/graph", getGraph)
		v1.GET("/area", getArea)
		v1.GET("/animals", getAnimals)
	}
	v2 := r.Group("/api/v2")
	{
		v2.GET("/animals/:name", func(ctx *gin.Context) {})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
