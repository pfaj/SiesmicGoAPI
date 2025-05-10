package services

// Project services houses the business logic for all project manipulation

import (
	"A5API/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

func Favicon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
}

func GetProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ProjectList)
}

func GetProjectByID(c *gin.Context) {
	var found bool = false
	// getting the id parameter from the url
	var str_id string = c.Param("id")

	// Converting to int 64 since the id is brought in as string
	id, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		// Force the function to leave early
		return
	}

	for _, project := range ProjectList {
		if project.ID == id {
			found = true
			c.IndentedJSON(http.StatusOK, project)
		}
	}

	if !found {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
	}
}

func GetProjectByCategory(c *gin.Context) {
	// getting the category parameter from the url
	var cat string = c.Param("category")

	// need to return a list of projects
	var projects = []models.Project{}

	for _, project := range ProjectList {
		if project.Category == cat {
			projects = append(projects, project)
		}
	}
	c.IndentedJSON(http.StatusOK, projects)
}

func UploadProject(c *gin.Context) {
	// Need to cast id to int64
	var id int64 = int64(len(ProjectList) + 1)
	var projectTitle string = c.PostForm("projectTitle")
	var clientName string = c.PostForm("clientName")
	var category string = c.PostForm("category")
	var description string = c.PostForm("description")
	var projectLink string = c.PostForm("projectLink")

	// initialize the projectstills
	var projectStills []string = []string{}

	// Intake Multipart form
	form, err := c.MultipartForm()

	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	files := form.File["projectStills"]

	for _, file := range files {
		var filename string = filepath.Base(file.Filename)

		projectStills = append(projectStills, filename)

		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	project := models.Project{ID: id,
		ProjectTitle:  projectTitle,
		ClientName:    clientName,
		Category:      category,
		ProjectStills: projectStills,
		Description:   description,
		ProjectLink:   projectLink}

	AddProject(project)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Form Submitted Succesfully."})

}

func EditProject(c *gin.Context) {
	var idStr string = c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var existingProject *models.Project
	var projectIndex int = -1
	//Search for project in project list
	for i, p := range ProjectList {
		if p.ID == id {
			// Make a copy of the current project that is stored
			existingProject = &p
			projectIndex = i
			break
		}
	}

	//Project does not exist
	if existingProject == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	projectTitle := c.PostForm("projectTitle")
	clientName := c.PostForm("clientName")
	category := c.PostForm("category")
	description := c.PostForm("description")
	projectLink := c.PostForm("projectLink")

	// Initialize the project stills for potential updates
	projectStills := existingProject.ProjectStills

	// Intake Multipart form
	form, err := c.MultipartForm()
	if err == nil {
		files := form.File["projectStills"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)

			projectStills = append(projectStills, filename)

			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
		}
	}

	if projectTitle != "" {
		existingProject.ProjectTitle = projectTitle
	}
	if clientName != "" {
		existingProject.ClientName = clientName
	}
	if category != "" {
		existingProject.Category = category
	}
	if description != "" {
		existingProject.Description = description
	}
	if projectLink != "" {
		existingProject.ProjectLink = projectLink
	}
	existingProject.ProjectStills = projectStills

	if projectIndex != -1 {
		ProjectList[projectIndex] = *existingProject
		writeProjectFile()
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Project with updated."})
}

func DeleteProject(c *gin.Context) {
	var found bool = false

	var str_id string = c.Param("id")

	id, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		// Force the function to leave early
		return
	}

	var i int = 0
	for _, project := range ProjectList {
		if project.ID == id {
			found = true
			RemoveProject(i)
		}
		i++
	}

	if !found {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Project deleted successfully."})
	}

}
