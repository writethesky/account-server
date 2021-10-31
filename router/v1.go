package router

import (
	v1 "account-server/api/v1"

	"github.com/gin-gonic/gin"
)

func version1(group *gin.RouterGroup) {
	tokenAPI := new(v1.Token)
	tokens := group.Group("tokens")
	tokens.POST("", tokenAPI.Create)

	userAPI := new(v1.User)
	users := group.Group("users")
	users.POST("", userAPI.Create)

}
