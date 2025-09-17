package main

import (
	"main/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", routes.Login)
		authGroup.POST("/logout", routes.Logout)
		authGroup.POST("/register", routes.Register)
	}

	r.Run(":8080")
}
