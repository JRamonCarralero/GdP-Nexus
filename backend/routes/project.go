package routes

import (
	"main/config"
	"main/controllers"
	"main/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	products, err := controllers.GetProjects(config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func CreateProject(c *gin.Context) {
	var req types.ProjectRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	projectId, err := controllers.CreateProject(config.DB, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = controllers.CreateNumberTaskIssue(config.DB, projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"projectId": projectId})
}

func UpdateProject(c *gin.Context) {
	var req types.ProjectUpdateRequest
	projectId := c.Param("id")

	if projectId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Project ID is required",
		})
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := controllers.UpdateProject(config.DB, projectId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully!"})
}

func DeleteProject(c *gin.Context) {
	projectId := c.Param("id")

	err := controllers.DeleteProject(config.DB, projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully!"})
}
