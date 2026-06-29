package main

import (
	"com.github/FelipecgPereira/go-jobs/db"
	"com.github/FelipecgPereira/go-jobs/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	server := gin.Default()
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	routes.RegisterRoutes(server)
	server.Run(":3000")
}
