package handler

import (
	"example/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertData(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	_, err := internal.MongoCollection.InsertOne(c, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
}

func DownloadExcel(c *gin.Context) {
	filename := "data.xlsx"
	if err := internal.ExportDataToExcel(filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to export data"})
		return
	}
	c.File(filename)
}
