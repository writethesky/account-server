package api

import (
	"account-server/internal"
	tokenV1 "account-server/pb/basic/token/v1"
	userV1 "account-server/pb/basic/user/v1"
	"net/http"
	"strconv"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/gin-gonic/gin"
)

type User struct {
}

type CreateUserRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"123456"`
}

// Create godoc
// @Summary add new users
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param userInfo body CreateUserRequest true "user info"
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
// @Summary Modifying User Information
// @Schemes
// @Tags users
// @Accept json
// @Produce json
// @Param _ body CreateUserRequest true "user info"
// @Param id path int true "user id. 0: current user"
// @Success 200 "Ok"
// @Failure 422 {object} Message
// @Router /users/{id} [patch]
func (*User) Patch(c *gin.Context) {
	// get router param: id
	idString := c.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if nil != err {
		c.JSON(http.StatusUnprocessableEntity, Message{err.Error()})
		return
	}

	var input CreateUserRequest
	if err := c.Bind(&input); nil != err {
		c.JSON(http.StatusUnprocessableEntity, Message{err.Error()})
		return
	}

	// todo Administrators can modify other user info
	auth, _ := c.Get("auth")
	if 0 != id && auth.(*tokenV1.ParseResponse).UserId != id {
		c.JSON(http.StatusForbidden, "Insufficient permissions")
		return
	}

	if 0 == id {
		id = auth.(*tokenV1.ParseResponse).UserId
	}

	if input.Password != "" {
		// modify password
		_, err = internal.NewUserServiceClient().SetPassword(c, &userV1.SetPasswordRequest{
			UserId:   id,
			Password: input.Password,
		})

		if nil != err {
			c.JSON(http.StatusUnprocessableEntity, status.Convert(err).Message())
			return
		}
	}
	c.Status(http.StatusOK)
}
