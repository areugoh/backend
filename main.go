package main

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := persistence.NewInMemoryStore(time.Minute * 60 * 24)

	v1 := r.Group("/api/v1")
	v1.Use(locationMiddleware())
	{
		v1.GET("/graph", getGraph)
		v1.GET("/area", getArea)
		v1.GET("/animals", cache.CachePage(store, time.Minute*60*24, getAnimals))
	}
	v2 := r.Group("/api/v2")
	{
		v2.GET("/animals/:name", func(ctx *gin.Context) {})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
