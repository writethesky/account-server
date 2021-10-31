package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	version1(v1)

	return r
}
