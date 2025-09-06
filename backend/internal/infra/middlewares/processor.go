package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HasProcessorAuthorization(processorSecret string, processorUserID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		processorPwd := c.GetHeader("Authorization")
		if processorPwd == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(processorSecret), []byte(processorPwd))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			return
		}

		c.Set("user_id", processorUserID)
		c.Next()
	}
}
