package middleware

import (
	"net/http"

	"github.com/PoteeDev/auth/auth"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
			c.Abort()
			return
		}
		metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
			return
		}
		if metadata.UserId != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "no access"})
			c.Abort()
		}
		c.Next()
	}
}
