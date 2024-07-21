package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorized(requireRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != requireRole {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "You are not authorized",
			})
			c.Abort()
		}
		c.Next()
	}
}