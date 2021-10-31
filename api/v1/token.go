package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Token struct {
}

type Message struct {
	Message string `json:"message"` //错误描述信息
}

type NotBody string

type CreateTokenRequest struct {
	Username string `json:"username" example:"admin"`  // 用户名
	Password string `json:"password" example:"123456"` // 密码
}

type CreateTokenResponse struct {
	Token  string `json:"token" example:"sadfasdfsadfdas"`
	Expire int64  `json:"expire" example:"1635646120"`
}

// Create godoc
// @Summary create token
// @Schemes
// @Tags tokens
// @Accept json
// @Produce json
// @Param userInfo body CreateTokenRequest true "用户信息"
// @Success 201 {object} CreateTokenResponse
// @Failure 401 {object} Message
// @Router /tokens [post]
func (*Token) Create(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Message{"用户名或密码错误"})
}
