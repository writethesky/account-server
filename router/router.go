package router

import (
	"account-server/api"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tokenAPI := new(api.Token)
	tokens := r.Group("tokens")
	tokens.POST("", tokenAPI.Create)

	userAPI := new(api.User)
	users := r.Group("users")
	users.POST("", userAPI.Create)

	return r
}
