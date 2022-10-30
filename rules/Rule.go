package rules

import (
	"github.com/gin-gonic/gin"
	"jwt-authentication/models"
)

type Rule interface {
	Run(user *models.User, context *gin.Context) (bool, string)
	NextRule() Rule
}
