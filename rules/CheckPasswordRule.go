package rules

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"jwt-authentication/models"
)

type CheckPasswordRule struct {
	Rule
}

func (c CheckPasswordRule) Run(user *models.User, context *gin.Context) (bool, string) {
	password := context.PostForm("password")

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return false, "Password is invalid!"
	}
	return true, ""
}

func (c CheckPasswordRule) NextRule() Rule {
	return CreateToken{}
}
