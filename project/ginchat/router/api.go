package router

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.GetIndex)

	return r
}
