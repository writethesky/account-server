package api

import (
	tokenV1 "account-server/pb/basic/token/v1"
	"account-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Account struct {
}

//
// type AccountNormal struct {
// 	Account  string `json:"account" example:"333333333"`
// 	Password string `json:"password" example:"12345678"`
// }
// type AccountEmail struct {
// 	Account      string `json:"account"`
// 	Password     string `json:"password"`
// 	LoginAddress string `json:"login_address"`
// }
//
// // List godoc
// // @Security ApiKeyAuth
// // @Summary 获取账号列表
// // @Description data有多种形式, 具体形式与type有关
// // @Description AccountTypeNormal
// // @Description <pre><code> {
// // @Description &nbsp;&nbsp;"account": "",
// // @Description &nbsp;&nbsp;"password": ""
// // @Description }</code></pre>
// // @Description AccountTypeEmail
// // @Description <pre><code> {
// // @Description &nbsp;&nbsp;"account": "",
// // @Description &nbsp;&nbsp;"password": "",
// // @Description &nbsp;&nbsp;"login_address": ""
// // @Description }</code></pre>
// // @Description AccountTypeSSH
// // @Description <pre><code> {
// // @Description &nbsp;&nbsp;"user": "",
// // @Description &nbsp;&nbsp;"password": "",
// // @Description &nbsp;&nbsp;"address": "",
// // @Description &nbsp;&nbsp;"port": 22
// // @Description }</code></pre>
// // @Tags accounts
// // @Accept json
// // @Produce json
// // @Success 200 {array} entity.Account{data=entity.AccountNormal}
// // @Router /accounts [get]
// func (*Account) List(c *gin.Context) {
//
// }

// Create godoc
// @Security ApiKeyAuth
// @Summary 新增账号
// @Tags accounts
// @Accept json
// @Produce json
// @Param _ body service.CreateAccountInput{data=entity.AccountNormal} true "账号信息"
// @Success 201 {object} model.Account{data=entity.AccountNormal}
// @Failure 422 {object} Message
// @Router /accounts/ [post]
func (*Account) Create(c *gin.Context) {
	params := new(service.CreateAccountInput)
	if err := c.Bind(params); nil != err {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	auth, _ := c.Get("auth")
	account, err := service.CreateAccount(auth.(*tokenV1.ParseResponse).UserId, *params)
	if nil != err {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	c.JSON(http.StatusCreated, account)
}

// Info godoc
// @Security ApiKeyAuth
// @Summary 获取账号详情
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "账号id"
// @Success 200 {object} model.Account{data=entity.AccountNormal}
// @Failure 404 {object} Message
// @Router /accounts/{id} [get]
func (*Account) Info(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if nil != err {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	auth, _ := c.Get("auth")
	account, err := service.GetAccountInfo(uint64(auth.(*tokenV1.ParseResponse).UserId), uint(id))
	if nil != err {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, account)
}

//
// // Delete godoc
// // @Security ApiKeyAuth
// // @Summary 删除账号
// // @Tags accounts
// // @Accept json
// // @Produce json
// // @Param id path string true "账号id"
// // @Success 204 {object} Message
// // @Failure 403 {object} Message
// // @Router /accounts/{id} [delete]
// func (*Account) Delete(c *gin.Context) {
//
// }
//
// // Put godoc
// // @Security ApiKeyAuth
// // @Summary 修改账号
// // @Tags accounts
// // @Accept json
// // @Produce json
// // @Param id path string true "账号id"
// // @Param _ body CreateAccountRequest{data=AccountNormal} true "账号信息"
// // @Success 200 {object} entity.Account{data=AccountNormal}
// // @Failure 422 {object} Message
// // @Router /accounts/{id} [put]
// func (*Account) Put(c *gin.Context) {
//
// }
