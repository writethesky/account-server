package api

import (
	"account-server/internal"
	userV1 "account-server/pb/basic/user/v1"
	"net/http"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
)

type User struct {
}

type CreateUserRequest struct {
	Username string `json:"username" example:"admin"`  // 用户名
	Password string `json:"password" example:"123456"` // 密码
}

// Create godoc
// @Summary 新增用户
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param userInfo body CreateUserRequest true "用户信息"
// @Success 201 "Created"
// @Failure 422 {object} Message
// @Router /users [post]
func (*User) Create(c *gin.Context) {
	var input CreateUserRequest
	if err := c.Bind(&input); nil != err {
		c.JSON(http.StatusUnprocessableEntity, Message{err.Error()})
		return
	}

	_, err := internal.NewUserServiceClient().Register(c, &userV1.RegisterRequest{
		Username: input.Username,
		Password: input.Password,
	})
	if nil != err && status.Code(err) == codes.Unavailable {
		c.JSON(http.StatusServiceUnavailable, Message{status.Convert(err).Message()})
		return
	}
	if nil != err {
		c.JSON(http.StatusUnprocessableEntity, Message{status.Convert(err).Message()})
		return
	}
	c.Status(http.StatusCreated)
}

// Patch godoc
// @Security ApiKeyAuth
// @Summary 修改用户信息
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param _ body service.CreateUserInput true "用户信息"
// @Success 200 "Ok"
// @Failure 422 {object} Message
// @Router /users [patch]
func (*User) Patch(c *gin.Context) {
}
