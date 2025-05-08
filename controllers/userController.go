package controllers

import (
	"A5API/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthUser(c *gin.Context) {
	services.AuthUser(c)
}

func VerifyToken(token string) (*jwt.Token, bool) {
	return services.VerifyToken(token)
}
