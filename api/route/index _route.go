package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewIndexRoute(group *gin.RouterGroup) {

	//todo  index.html
	group.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "You reached the album management site...",
		})
	})
}
