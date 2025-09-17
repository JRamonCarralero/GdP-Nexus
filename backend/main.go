package main

import (
	"log"
	"net/http"

	"main/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Nexus API
// @version 1.0
// @description API para el proyecto Nexus
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error: No se encontró el archivo .env, usando variables de entorno del sistema.")
	}

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
