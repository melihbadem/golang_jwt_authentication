package routes

import (
	"github.com/gin-gonic/gin"
	c "jwt-authentication/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/login", new(c.AuthController).Login)
	router.POST("/register", new(c.AuthController).Register)
	router.POST("/logout", new(c.AuthController).Logout)
	router.GET("/user", new(c.UserController).User)
}
