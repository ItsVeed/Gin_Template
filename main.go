package main

import (
	"github.com/ItsVeed/Gin_Template/controllers"
	"github.com/ItsVeed/Gin_Template/initializers"
	"github.com/ItsVeed/Gin_Template/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
