package routes

//Making a group of routes here so that I can keep everything in their own files making it more organized

import (
	"A5API/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

var router = gin.Default()
var limiter = rate.NewLimiter(1, 5)

func Run() {
	// Adding a rete limiter to the entire API so that it cannot be over loaded
	router.Use(rateLimiter)
	router.Use(applyCors)

	// Get All router groups
	getRouters()

	router.GET("/fav.ico", controllers.Favicon)

	// Run API on port 9010
	router.Run(":9010")
}

func getRouters() {
	// Added routes to v1 extenstion
	// Ex. /v1/projects
	v1 := router.Group("/v1")
	addProjectRoutes(v1)
	addUserRoutes(v1)
}

// Rate Limits the amount of requests from a user
func rateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		fmt.Println("Rate Limit")
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
		c.Abort()
		return
	}
	c.Next()
}

func authenticateMiddleware(c *gin.Context) {
	// Retrieve the token from the cookie
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		fmt.Println("Token missing in header")
		c.Abort()
		return
	}
	// Verify the token
	token, valid := controllers.VerifyToken(tokenString[6:])
	if !valid {
		fmt.Println("Token not valid")
		c.Abort()
		return
	}
	// Print information about the verified token
	fmt.Println("Token verified successfully. Claims: ", token.Claims)
	// Continue with the next middleware or route handler
	c.Next()
}
func applyCors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}
