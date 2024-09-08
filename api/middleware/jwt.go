package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/domain"
	"github.com/hxzhouh/go-zen.git/internal"
	"net/http"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
			c.Abort()
		}
		authorized, err := internal.IsAuthorized(authToken, secret)
		if authorized {
			userID, err := internal.ExtractIDFromToken(authToken, secret)
			if err != nil {
				c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
				c.Abort()
				return
			}
			c.Set("x-user-id", userID)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
		c.Abort()
		return
	}
}
