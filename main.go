package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"slackUsername": "Thoth",
			"backend":       true,
			"age":           22,
			"bio":           "i am a vey hardworking engineer that loves basketball",
		})
	})

	r.Run()
}
