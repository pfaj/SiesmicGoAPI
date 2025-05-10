package main

import (
	"A5API/routes"
)

func main() {
	//Calling the Run on the routes package where I can then handle all the
	//routes and grouping of different parts of API calls
	routes.Run()
}

// References
// https://go.dev/doc/
// https://go.dev/doc/effective_go
// https://github.com/gin-gonic/examples
// https://golang-jwt.github.io/jwt/usage/create/#with-default-options
