package controller

import (
	http "net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "create user"})
}

func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "get all users"})
}

func UpdateUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "update user"})
}

func DeleteUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "delete user"})
}
