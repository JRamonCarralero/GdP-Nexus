package routes

import (
	"main/config"
	"main/controllers"
	"main/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all projects
// @Description Returns a list of projects
// @Tags Project
// @Accept json
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {object} object "Internal server error"
// @Router /projects [get]
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

// @Summary Get a project by ID
// @Description Returns a project by ID
// @Tags Project
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} models.Project
// @Failure 404 {object} object "Not found"
// @Failure 500 {object} object "Internal server error"
// @Router /projects/{id} [get]
func GetProject(c *gin.Context) {
	projectId := c.Param("id")

	project, err := controllers.GetProjectById(config.DB, projectId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, project)
}

// @Summary Create a new project
// @Description Creates a new project
// @Tags Project
// @Accept json
// @Produce json
// @Param request body types.ProjectRequest true "Project request"
// @Success 201 {object} object "Project ID"
// @Failure 400 {object} object "Bad request"
// @Failure 500 {object} object "Internal server error"
// @Router /projects [post]
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

// @Summary Update a project
// @Description Updates a project by ID
// @Tags Project
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param request body types.ProjectUpdateRequest true "Project update request"
// @Success 200 {object} object "Project updated successfully!"
// @Failure 400 {object} object "Bad request"
// @Failure 500 {object} object "Internal server error"
// @Router /projects/{id} [put]
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

// @Summary Delete a project
// @Description Deletes a project by ID
// @Tags Project
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} object "Project deleted successfully!"
// @Failure 500 {object} object "Internal server error"
// @Router /projects/{id} [delete]
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
