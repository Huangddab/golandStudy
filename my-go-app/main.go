package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func main() {
	r := gin.Default()
	r.GET("/export", func(c *gin.Context) {
		file := excelize.NewFile()
		sheetName := "Sheet1"
		// 向工作表中写入数据
		for row := 1; row <= 10; row++ {
			for col := 1; col <= 5; col++ {
				cell, err := excelize.CoordinatesToCellName(col, row)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cell name"})
					return
				}
				if err := file.SetCellValue(sheetName, cell, fmt.Sprintf("Row %d, Col %d", row, col)); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cell value"})
					return
				}
			}
		}
		// 设置 HTTP 响应的头信息
		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		filename := fmt.Sprintf("Report_%v.xlsx", time.Now().Format("2006-01-02T15:04:05"))
		c.Header("Content-Disposition", "attachment; filename="+filename)
		// 将 Excel 文件写入 HTTP 响应
		if err := file.Write(c.Writer); err != nil {
			c.JSON(http.StatusOK, "failed")
			return
		}
		c.JSON(http.StatusOK, "success")
	})
	r.Run(":8080")
}
