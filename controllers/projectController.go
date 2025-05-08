package controllers

// Project Controllers pointing to which projects service functions to call

import (
	"A5API/services"
	"github.com/gin-gonic/gin"
)

func Favicon(c *gin.Context) {
	services.Favicon(c)
}

func GetProjects(c *gin.Context) {
	services.GetProjects(c)
}

func GetProjectByID(c *gin.Context) {
	services.GetProjectByID(c)
}

func GetProjectByCategory(c *gin.Context) {
	services.GetProjectByCategory(c)
}

func UploadProject(c *gin.Context) {
	services.UploadProject(c)
}

func EditProject(c *gin.Context) {
	services.EditProject(c)
}

func DeleteProject(c *gin.Context) {
	services.DeleteProject(c)
}
