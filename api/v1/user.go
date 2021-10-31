package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

type CreateUserResponse struct {
}

// Create godoc
// @Summary 新增用户
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param userInfo body CreateTokenRequest true "用户信息"
// @Success 201 "Created"
// @Failure 422 {object} Message
// @Router /users [post]
func (*User) Create(c *gin.Context) {
	c.Status(http.StatusCreated)
}

// Patch godoc
// @Security ApiKeyAuth
// @Summary 修改用户信息
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param _ body CreateTokenRequest true "用户信息"
// @Success 200 "Ok"
// @Failure 422 {object} Message
// @Router /users [patch]
func (*User) Patch(c *gin.Context) {
}
