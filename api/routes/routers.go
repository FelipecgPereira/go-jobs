package routes

import (
	"com.github/FelipecgPereira/go-jobs/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/auth", auth)

	customerGroup := server.Group("/customer")
	customerGroup.Use(middlewares.Autheticate)
	customerGroup.POST("/", createCustomer)
	customerGroup.PUT("/:id", updateCustomer)
	customerGroup.GET("/", getCustomers)
	customerGroup.GET("/:id", getCustmerById)

}
