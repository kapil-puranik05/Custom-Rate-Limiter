package main

import (
	"rate_limiter/internal/database"
	"rate_limiter/internal/initializer"
	"rate_limiter/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	initializer.InitializeEnvironmentVariables()
	database.ConnectCache()
	r := gin.Default()
	routes.GetApiRoutes(r)
	r.Run(":8080")
}
