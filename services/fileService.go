package services

import (
	"A5API/models"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
)

// Public Variables
var ProjectList []models.Project
var UserList []models.User

// Initialize the project and users
func init() {
	openProjectFile()
	openUserFile()
}

// PUBLIC FUNCTIONS //

func AddProject(project models.Project) {
	ProjectList = append(ProjectList, project)

	writeProjectFile()
}

func RemoveProject(id int) {
	ProjectList = slices.Delete(ProjectList, id, id+1)

	writeProjectFile()
}

// PRIVATE FUNCTIONS //
func openUserFile() {
	jsonFile, err := os.Open("files/users.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")

	UserList = parseUserJSON(jsonFile)

	// defer the closing of the file
	defer jsonFile.Close()
}

func openProjectFile() {
	jsonFile, err := os.Open("files/projects.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened projects.json")

	ProjectList = parseProjectJSON(jsonFile)

	// defer the closing file
	defer jsonFile.Close()

}

func parseProjectJSON(file *os.File) []models.Project {
	var list = []models.Project{}
	byteValue, err := io.ReadAll(file)

	if err != nil {
		println("Error reading file:", err.Error())
		return list
	}

	json.Unmarshal(byteValue, &list)

	//TEST: for loop to make sure all projects were read.
	// for i := 0; i < len(list); i++ {
	// 	fmt.Println(fmt.Sprintf("parseJSON list.Project: %v", list[i].ProjectTitle))
	// }

	return list

}

func parseUserJSON(file *os.File) []models.User {
	var list = []models.User{}
	byteValue, err := io.ReadAll(file)

	if err != nil {
		println("Error reading file:", err.Error())
		return list
	}

	json.Unmarshal(byteValue, &list)

	return list

}

func writeProjectFile() {
	JSON, err := json.Marshal(ProjectList)
	if err != nil {
		println("Error marshaling JSON:", err.Error())
		return
	}
	err = os.WriteFile("./files/projects.json", JSON, 0644)
	if err != nil {
		println("Error writing file:", err.Error())
		return
	}
}
