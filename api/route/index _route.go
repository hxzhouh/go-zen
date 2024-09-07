package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewIndexRoute(group *gin.RouterGroup) {

	//todo  index.html
	group.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "You reached the album management site...",
		})
	})
	//group.GET("/home", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "home.html", gin.H{
	//		"content": "You reached the album management site...",
	//	})
	//})
}
