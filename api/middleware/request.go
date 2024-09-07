package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func RequestLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		request_id := c.Request.Header.Get("request_id")
		slog.Info(fmt.Sprintf("Request Method: %s %s, Request ID: %s", c.Request.Method, c.Request.URL.Path, request_id))
		c.Next()
		return
	}
}
