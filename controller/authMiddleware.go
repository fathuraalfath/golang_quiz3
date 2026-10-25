package controller

import (
	"golang_quiz3/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthMiddleware implements HTTP Basic Authentication against the database users table
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Basic auth required"})
			return
		}

		if database.DB == nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "database not connected"})
			return
		}

		var user database.User
		if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid username or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid username or password"})
			return
		}

		c.Set("user_id", int(user.ID))
		c.Next()
	}
}
