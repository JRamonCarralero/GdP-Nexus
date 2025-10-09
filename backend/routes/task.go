package routes

import (
	"main/internal/config"
	"main/internal/controllers"
	"main/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all tasks
// @Description Returns a list of tasks
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 500 {object} object "Internal server error"
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	projects, err := controllers.GetProjects(config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// @Summary Get a task by ID
// @Description Returns a task by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} models.Task
// @Failure 404 {object} object "Not found"
// @Failure 500 {object} object "Internal server error"
// @Router /tasks/{id} [get]
func GetTask(c *gin.Context) {
	taskId := c.Param("id")

	task, err := controllers.GetTaskById(config.DB, taskId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

// @Summary Create a new task
// @Description Creates a new task
// @Tags Task
// @Accept json
// @Produce json
// @Param request body types.TaskRequest true "Task request"
// @Success 201 {object} object "Task ID"
// @Failure 400 {object} object "Bad request"
// @Failure 500 {object} object "Internal server error"
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var req types.TaskRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	projectId, err := controllers.CreateTask(config.DB, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, projectId)
}

// @Summary Update a task
// @Description Updates a task by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param request body types.TaskUpdateRequest true "Task update request"
// @Success 200 {object} object "Task updated successfully!"
// @Failure 400 {object} object "Bad request"
// @Failure 500 {object} object "Internal server error"
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	var req types.TaskUpdateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := controllers.UpdateTask(config.DB, taskId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})
}

// @Summary Delete a task
// @Description Deletes a task by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} object "Task deleted successfully!"
// @Failure 500 {object} object "Internal server error"
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")

	err := controllers.DeleteTask(config.DB, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})
}
