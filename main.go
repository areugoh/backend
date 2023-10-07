package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/v1/graph/:long/:lat", getGraph)
	r.GET("/api/v1/area/:long/:lat", getArea)
	r.GET("/api/v1/animals/:long/:lat", getAnimals)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
