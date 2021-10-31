# 账号管家 服务端


## 生成API文档

1. 安装swag工具，若已安装请忽略
`go get -u github.com/swaggo/swag/cmd/swag`
2. 生成文档 `swag init --instanceName v1 --output ./docs/v1 --dir ./api/v1 -g _doc.go`
3. 访问文档 http://localhost:8080/swagger/v1/index.html