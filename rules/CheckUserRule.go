package rules

import (
	"github.com/gin-gonic/gin"
	"jwt-authentication/database"
	"jwt-authentication/models"
)

type CheckUserRule struct {
	Rule
}

func (c CheckUserRule) Run(user *models.User, context *gin.Context) (bool, string) {
	email := context.PostForm("email")

	database.DBConn.Where("email = ?", email).First(&user)

	if user.Id == 0 {
		return false, "User not found!"
	}
	return true, ""
}

func (c CheckUserRule) NextRule() Rule {
	return CheckPasswordRule{}
}
