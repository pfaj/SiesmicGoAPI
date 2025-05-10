package routes

// All project routes that are used in the API

import (
	"A5API/controllers"
	"github.com/gin-gonic/gin"
)

func addProjectRoutes(rg *gin.RouterGroup) {
	var projects *gin.RouterGroup = rg.Group("/projects")
	projects.GET("/", controllers.GetProjects)
	projects.PUT("/:id", authenticateMiddleware, controllers.EditProject)
	projects.GET(":id", controllers.GetProjectByID)
	projects.POST("/upload", authenticateMiddleware, controllers.UploadProject)
	projects.DELETE("/:id", authenticateMiddleware, controllers.DeleteProject)
	// projects.GET("/:category", controllers.GetProjectByCategory)
}
