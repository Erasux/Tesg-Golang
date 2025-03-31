package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/events")
	router.GET("/events")
	router.GET("/events/:id")
	router.PUT("/events/:id")
	router.DELETE("/events/:id")
	router.Run(":8080")
}
