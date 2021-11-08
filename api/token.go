package api

import (
	"account-server/internal"
	tokenV1 "account-server/pb/basic/token/v1"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Token struct {
}

type Message struct {
	Message string `json:"message"` // 错误描述信息
}

type CreateTokenRequest struct {
	Username string `json:"username" example:"user"`   // 用户名
	Password string `json:"password" example:"123456"` // 密码
}

type CreateTokenResponse struct {
	Token  string    `json:"token" example:"sadfasdfsadfdas"`
	Expire time.Time `json:"expire" example:"1635646120"`
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
	var input CreateTokenRequest
	if err := c.Bind(&input); nil != err {
		c.JSON(http.StatusUnprocessableEntity, Message{err.Error()})
		return
	}
	res, err := internal.NewTokenServiceClient().Apply(c, &tokenV1.ApplyRequest{
		Username: input.Username,
		Password: input.Password,
	})
	if nil != err {
		c.JSON(http.StatusUnauthorized, Message{"The username or password is incorrect"})
		return
	}

	c.JSON(http.StatusCreated, CreateTokenResponse{res.Token, res.Expire.AsTime()})
}
