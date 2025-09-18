package main

import (
	"context"
	"log"
	"net/http"

	"main/config"
	"main/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Nexus API
// @version 1.0
// @description API for Nexus Project
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: .env file not found")
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Fatal error connecting to database: %v", err)
	}
	defer config.DB.Disconnect(context.TODO())

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", routes.Login)
		authGroup.POST("/register", routes.Register)
	}

	r.Run(":8080")
}
