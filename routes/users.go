package routes

// All project routes that are used in the API

import (
	"A5API/controllers"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	var auth *gin.RouterGroup = rg.Group("/auth")
	auth.POST("/", controllers.AuthUser)
}
