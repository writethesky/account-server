package api

import "github.com/gin-gonic/gin"

type Account struct {
}

type AccountType uint

const (
	AccountTypeNormal AccountType = iota
	AccountTypeEmail
	AccountTypeSSH
)

type AccountNormal struct {
	Account  string `json:"account" example:"333333333"`
	Password string `json:"password" example:"12345678"`
}
type AccountEmail struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	LoginAddress string `json:"login_address"`
}
type AccountEntity struct {
	ID    int64       `json:"id"`
	Title string      `json:"title" example:"常用QQ"`
	Type  AccountType `json:"type"` // Enums(AccountTypeNormal, AccountTypeEmail, AccountTypeSSH)
	Data  interface{} `json:"data"`
}

// List godoc
// @Security ApiKeyAuth
// @Summary 获取账号列表
// @Description data有多种形式, 具体形式与type有关
// @Description AccountTypeNormal
// @Description <pre><code> {
// @Description &nbsp;&nbsp;"account": "",
// @Description &nbsp;&nbsp;"password": ""
// @Description }</code></pre>
// @Description AccountTypeEmail
// @Description <pre><code> {
// @Description &nbsp;&nbsp;"account": "",
// @Description &nbsp;&nbsp;"password": "",
// @Description &nbsp;&nbsp;"login_address": ""
// @Description }</code></pre>
// @Description AccountTypeSSH
// @Description <pre><code> {
// @Description &nbsp;&nbsp;"user": "",
// @Description &nbsp;&nbsp;"password": "",
// @Description &nbsp;&nbsp;"address": "",
// @Description &nbsp;&nbsp;"port": 22
// @Description }</code></pre>
// @Tags accounts
// @Accept json
// @Produce json
// @Success 200 {array} AccountEntity{data=AccountNormal}
// @Router /accounts [get]
func (*Account) List(c *gin.Context) {

}

type CreateAccountRequest struct {
	Title string      `json:"title" example:"常用QQ"`
	Type  AccountType `json:"type"` // Enums(AccountTypeNormal, AccountTypeEmail, AccountTypeSSH)
	Data  interface{} `json:"data"`
}

// Create godoc
// @Security ApiKeyAuth
// @Summary 新增账号
// @Tags accounts
// @Accept json
// @Produce json
// @Param _ body CreateAccountRequest{data=AccountNormal} true "账号信息"
// @Success 201 {object} AccountEntity{data=AccountNormal}
// @Failure 422 {object} Message
// @Router /accounts/ [post]
func (*Account) Create(c *gin.Context) {

}

// Info godoc
// @Security ApiKeyAuth
// @Summary 获取账号详情
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "账号id"
// @Success 200 {object} AccountEntity{data=AccountNormal}
// @Failure 404 {object} Message
// @Router /accounts/{id} [get]
func (*Account) Info(c *gin.Context) {

}

// Delete godoc
// @Security ApiKeyAuth
// @Summary 删除账号
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "账号id"
// @Success 204 {object} Message
// @Failure 403 {object} Message
// @Router /accounts/{id} [delete]
func (*Account) Delete(c *gin.Context) {

}

// Put godoc
// @Security ApiKeyAuth
// @Summary 修改账号
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "账号id"
// @Param _ body CreateAccountRequest{data=AccountNormal} true "账号信息"
// @Success 200 {object} AccountEntity{data=AccountNormal}
// @Failure 422 {object} Message
// @Router /accounts/{id} [put]
func (*Account) Put(c *gin.Context) {

}
