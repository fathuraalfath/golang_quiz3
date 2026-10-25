package controller

import (
	http "net/http"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "create category"})
}

func GetAllCategory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "get all category"})
}

func UpdateCategory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "update category"})
}

func DeleteCategory(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "delete category"})
}
