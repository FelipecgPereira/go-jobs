package routes

import (
	"net/http"

	"com.github/FelipecgPereira/go-jobs/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request body"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}
