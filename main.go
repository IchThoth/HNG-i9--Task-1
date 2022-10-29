package main

import (
	"fmt"
	"log"

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
	err := r.Run(":3031")
	fmt.Println("started server at 3031")
	if err != nil {
		log.Fatal(err)
	}
}
