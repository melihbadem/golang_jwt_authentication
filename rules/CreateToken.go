package rules

import (
	"github.com/gin-gonic/gin"
	"jwt-authentication/models"
	"jwt-authentication/services"
	"os"
)

type CreateToken struct {
	Rule
}

func (c CreateToken) Run(user *models.User, _ *gin.Context) (bool, string) {
	var jwtService services.JwtService
	token, err := jwtService.CreateToken(user.Id, os.Getenv("SECRET_KEY"))

	if err != nil {
		return false, "User not logged in!"
	}
	return true, token
}

func (c CreateToken) NextRule() Rule {
	return nil
}
