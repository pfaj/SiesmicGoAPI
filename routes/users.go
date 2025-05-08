package routes

// All project routes that are used in the API

import (
	"A5API/controllers"
	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	projects := rg.Group("/auth")
	projects.POST("/", controllers.AuthUser)
}
