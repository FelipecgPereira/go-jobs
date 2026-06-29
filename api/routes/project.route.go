package routes

import (
	"net/http"
	"strconv"

	"com.github/FelipecgPereira/go-jobs/models"
	"github.com/gin-gonic/gin"
)

func createProject(context *gin.Context) {
	var project models.Project

	err := context.ShouldBindJSON(&project)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request body"})
		return
	}

	project.UserId = context.GetInt64("userId")

	id, err := project.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create project", "error": err})
		return
	}

	context.Header("Location", "/project/"+strconv.FormatInt(id, 10))
	context.Header("X-Id", strconv.FormatInt(id, 10))
	context.JSON(http.StatusCreated, gin.H{"message": "Project created successfully"})

}

func getProjects(context *gin.Context) {
	userId := context.GetInt64("userId")
	projects, err := models.GetProjects(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve projects"})
		return
	}

	context.JSON(http.StatusOK, projects)
}

func getProjectById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid project ID"})
		return
	}

	project, err := models.GetProjectById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve project"})
		return
	}
	if project == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "project not found"})
		return
	}
	context.JSON(http.StatusOK, project)
}

func updateProject(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid project ID"})
		return
	}

	userId := context.GetInt64("userId")
	project, err := models.GetProjectById(id)

	if project.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	var updateProject models.Project
	if err := context.ShouldBindJSON(&updateProject); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	updateProject.Id = id
	err = updateProject.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update project"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})

}
