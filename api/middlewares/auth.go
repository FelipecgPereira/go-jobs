package middlewares

import (
	"net/http"

	"com.github/FelipecgPereira/go-jobs/utils"
	"github.com/gin-gonic/gin"
)

func Autheticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not autorized",
		})
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
