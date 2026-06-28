package main

import (
	"com.github/FelipecgPereira/go-jobs/db"
	"com.github/FelipecgPereira/go-jobs/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":3000")
}
