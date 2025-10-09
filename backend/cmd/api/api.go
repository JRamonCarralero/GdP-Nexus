package api

import (
	"main/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunAPIServer() error {
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

	projectGroup := r.Group("/projects")
	{
		projectGroup.GET("/", routes.GetProjects)
		projectGroup.GET("/:id", routes.GetProject)
		projectGroup.POST("/", routes.CreateProject)
		projectGroup.PUT("/:id", routes.UpdateProject)
		projectGroup.DELETE("/:id", routes.DeleteProject)
	}

	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("/", routes.GetTasks)
		taskGroup.GET("/:id", routes.GetTask)
		taskGroup.POST("/", routes.CreateTask)
		taskGroup.PUT("/:id", routes.UpdateTask)
		taskGroup.DELETE("/:id", routes.DeleteTask)
	}

	r.Run(":8080")

	return nil
}
