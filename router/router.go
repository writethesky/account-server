package router

import (
	"account-server/api"
	"account-server/internal"
	tokenV1 "account-server/pb/basic/token/v1"
	"net/http"
	"strings"

	"google.golang.org/grpc/status"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("secures", cors, api.SetSecureKey)
	r.Use(cors, api.Secure)

	tokenAPI := new(api.Token)
	tokens := r.Group("tokens")
	tokens.POST("", tokenAPI.Create)

	userAPI := new(api.User)
	users := r.Group("users")
	users.POST("", userAPI.Create)
	users.PATCH("/:id", auth, userAPI.Patch)

	accountAPI := new(api.Account)
	accounts := r.Group("accounts")
	accounts.POST("", auth, accountAPI.Create)
	accounts.GET("", auth, accountAPI.List)
	accounts.DELETE("/:id", auth, accountAPI.Delete)
	accounts.GET("/:id", auth, accountAPI.Info)
	accounts.PUT("/:id", auth, accountAPI.Put)

	return r
}

func cors(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, Security, accept, origin, Cache-Control, Request-ID, Nonce, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
}

func auth(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	authorizationArr := strings.Split(authorization, " ")
	if 2 != len(authorizationArr) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization:AuthType AuthValue. e.g. Authorization:token xxxxxx")
		return
	}
	authorizationType := authorizationArr[0]
	authorizationValue := authorizationArr[1]
	if "token" != authorizationType {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Currently, only token are supported")
		return
	}

	// check token
	res, err := internal.NewTokenServiceClient().Parse(c, &tokenV1.ParseRequest{
		Token: authorizationValue,
	})
	if nil != err {
		c.AbortWithStatusJSON(http.StatusUnauthorized, status.Convert(err).Message())
		return
	}
	c.Set("auth", res)

}
