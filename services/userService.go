package services

import (
	"A5API/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func AuthUser(c *gin.Context) {
	var parsed models.User

	body, _ := io.ReadAll(c.Request.Body)

	// Turning the body into a struct so that the login input can be read
	json.Unmarshal(body, &parsed)

	for i := range UserList {
		if UserList[i].Username == parsed.Username {
			var user models.User = UserList[i]
			// Compares the input password the stored hashed password
			hashMatch := CheckPasswordHash(parsed.Password, user.Password)

			if hashMatch {
				// we will prepare a JSON Web Token to return to the user
				var jwt string = CreateJWT(user.Username)
				// c.SetCookie("Authentication", jwt, 3600, "/", "", false, true)
				c.IndentedJSON(http.StatusOK, gin.H{"message": "User Authorized", "jwt": jwt, "auth": true})
				// want to exit early so that we just leave one response instead of two
				return
			}
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Invalid username or password", "auth": false})
}
