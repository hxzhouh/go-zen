package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func NewSwaggerRouter(group *gin.RouterGroup) {
	group.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
}
