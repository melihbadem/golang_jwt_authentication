package controllers

import (
	"github.com/gin-gonic/gin"
	"jwt-authentication/models"
	"jwt-authentication/rules"
	"jwt-authentication/services"
	"net/http"
	"time"
)

type AuthController struct {
}

func (c AuthController) Register(context *gin.Context) {
	name := context.PostForm("name")
	email := context.PostForm("email")
	password := context.PostForm("password")

	userService := services.UserService{}
	user := userService.CreateUser(name, email, password)

	context.JSON(http.StatusOK, user)
}

func (c AuthController) Login(context *gin.Context) {
	var user models.User
	var rule rules.Rule = rules.CheckUserRule{}

	for rule.NextRule() != nil {
		if isOk, err := rule.Run(&user, context); isOk != true {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": err,
			})
			return
		}
		rule = rule.NextRule()
	}
	if rule.NextRule() == nil {
		if isOk, token := rule.Run(&user, context); isOk {
			services.SetCookie(context.Writer, "jwt", token, time.Hour*24)
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}
	}
}

func (c AuthController) Logout(context *gin.Context) {
	services.SetCookie(context.Writer, "jwt", "", -time.Hour)
	context.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
