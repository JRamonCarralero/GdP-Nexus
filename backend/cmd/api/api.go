package api

import (
	"main/routes"
	"net/http"

	_ "main/docs" // Importa los docs generados, ¡IMPORTANTE!

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RunAPIServer sets up a Gin router and runs an HTTP server on port 8080.
// The server handles the following endpoints:
// - GET /: returns a "Hello, World!" message.
// - POST /auth/login: handles user login.
// - POST /auth/register: handles user registration.
// - GET /projects: returns a list of projects.
// - GET /projects/:id: returns a project by ID.
// - POST /projects: creates a new project.
// - PUT /projects/:id: updates a project by ID.
// - DELETE /projects/:id: deletes a project by ID.
// - GET /tasks: returns a list of tasks.
// - GET /tasks/:id: returns a task by ID.
// - POST /tasks: creates a new task.
// - PUT /tasks/:id: updates a task by ID.
// - DELETE /tasks/:id: deletes a task by ID.
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")

	return nil
}
