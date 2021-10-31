package main

import (
	_ "account-server/docs/v1"
	"account-server/router"
	"account-server/swagger"
)

func main() {
	r := router.Init()

	//v1.SwaggerInfo.BasePath = "/v11"
	swagger.Run(r)

	r.Run(":8080")
}
