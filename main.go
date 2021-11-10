package main

import (
	_ "account-server/docs"
	"account-server/internal"
	"account-server/repository/entity"
	"account-server/router"
	"fmt"
	"log"
)

// @title Account API
// @version 1.0
// @description This is an account server.
// @description ## 错误分类
// @description * **客户端错误：**是不应该让后端感知到的错误，即`应该消灭在客户端`层面的错误。由于客户端请求参数不符合要求而产生的错误。包括但不限于参数格式错误、缺少参数、参数类型、参数长度等错误
// @description * **用户行为错误：**是应该让后端感知到的错误，因为`客户端无法消灭`该错误。由于用户行为有意或无意产生的错误。包括但不限于输入了错误的用户名或密码（没输入属于客户端错误中的缺少参数或参数长度错误）；创建资源时输入了系统中已经存在的资源导致无法创建。
// @description * **服务端错误：**是在正常情况下`不应该发生`的错误，属于服务端的故障导致的错误。
// @description ## 客户端错误
// @description 接收请求正文的 API 调用上可能存在以下种类型的客户端错误：
// @description ### 1 发送无效的 JSON 将导致 `400 Bad Request` 响应。
// @description ```
// @description HTTP/2 400
// @description Content-Length: 35
// @description
// @description {"message":"Problems parsing JSON"}
// @description ```
// @description ### 2 发送错误类型的 JSON 值将导致 `400 Bad Request` 响应。
// @description ```
// @description HTTP/2 400
// @description Content-Length: 40
// @description
// @description {"message":"Body should be a JSON object"}
// @description ```
// @description ### 3 发送无效的字段将导致 `422 Unprocessable Entity` 响应。
// @description <pre><code>
// @description HTTP/2 422
// @description Content-Length: 149<br>
// @description {
// @description &nbsp;&nbsp;"message": "Validation Failed",
// @description &nbsp;&nbsp;"errors": [
// @description &nbsp;&nbsp;&nbsp;&nbsp;{
// @description &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"resource": "Issue",
// @description &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"field": "title",
// @description &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"code": "missing_field"
// @description &nbsp;&nbsp;&nbsp;&nbsp;}
// @description &nbsp;&nbsp;]
// @description }
// @description </code></pre>
// @description
// @description ### 4 需要授权的接口没有正常传递token或传递了错误、过期的token `401 Unauthorized` 响应。
// @description ## 用户行为错误
// @description ### 1 申请token时传递了不正确的账号密码 `401 Unauthorized` 响应。
// @description ### 2 验证失败，如名称重复等 `422 Unprocessable Entity` 响应。
// @description ## 错误列出说明
// @description 下方仅列出接口可能产生的`用户行为错误`，不列出`客户端错误`

// @contact.name writethesky
// @contact.url https://github.com/writethesky

// @host localhost:8080
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

// @Tag.name accounts
// @Tag.description 账号相关，注意，具体data的结构与type有关，详细介绍见【GET /accounts】

func main() {
	migrate()

	r := router.Init()

	log.Fatalln(r.Run(fmt.Sprintf(":%d", internal.Config.Server.Port)))
}

func migrate() {
	err := internal.DB.AutoMigrate(new(entity.Account))
	if nil != err {
		panic(err)
	}
}
