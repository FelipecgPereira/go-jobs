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
	customerGroup.POST("", createCustomer)
	customerGroup.PUT("/:id", updateCustomer)
	customerGroup.GET("", getCustomers)
	customerGroup.GET("/:id", getCustmerById)

	projectGroup := server.Group("/project")
	projectGroup.Use(middlewares.Autheticate)
	projectGroup.POST("/", createProject)
	projectGroup.PUT("/:id", updateProject)
	projectGroup.GET("", getProjects)
	projectGroup.GET("/:id", getProjectById)

	b2bGroup := server.Group("/b2b")
	b2bGroup.Use(middlewares.Autheticate)
	b2bGroup.POST("/", createB2b)
	b2bGroup.PUT("/:id", updateB2b)
	b2bGroup.GET("/sum", sumPayments)

}
