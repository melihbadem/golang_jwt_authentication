package controllers

import (
	"github.com/gin-gonic/gin"
	"jwt-authentication/services"
	"net/http"
	"os"
)

type UserController struct {
}

func (c UserController) User(context *gin.Context) {
	getJwtToken, _ := context.Cookie("jwt")

	var jwtService services.JwtService
	token, err := jwtService.GetToken(getJwtToken, os.Getenv("SECRET_KEY"))

	if err != nil {
		context.JSON(http.StatusUnauthorized, err)
		return
	}
	claims := jwtService.GetClaims(token)

	userService := services.UserService{}
	user := userService.GetUser(claims.Issuer)
	context.JSON(http.StatusOK, user)
}
