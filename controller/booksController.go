package controller

import (
	http "net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "create book"})
}

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "get all books"})
}

func UpdateBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "update book"})
}

func DeleteBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "delete book"})
}
