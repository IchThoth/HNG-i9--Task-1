package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r.Run(":" + port)
}
