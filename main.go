package main

import (
	"A5API/controllers"
	"A5API/routes"
)

func main() {
	// Initialize the project list from a json file
	controllers.OpenFile()

	//Calling the Run on the routes package where I can then handle all the
	//routes and grouping of different parts of API calls
	routes.Run()
}
