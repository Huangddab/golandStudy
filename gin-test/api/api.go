package api

import (
	"example/api/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/insert", handler.InsertData)
	r.GET("/download", handler.DownloadExcel)
	return r
}
