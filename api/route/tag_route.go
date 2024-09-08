package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/api/controller"
	"github.com/hxzhouh/go-zen.git/storage"
	"github.com/hxzhouh/go-zen.git/storage/sqlite"
	"github.com/hxzhouh/go-zen.git/usecase"
	"time"
)

func NewTagRouter(timeout time.Duration, publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	pr := sqlite.NewTagRepository(storage.DefaultStorage)
	pc := &controller.TagController{
		TagUsecase: usecase.NewTagUsecase(pr, timeout),
		Env:        nil,
	}
	tagPublicGroup := publicGroup.Group("/tag")
	{
		tagPublicGroup.GET("/list", pc.List)
		tagPublicGroup.GET("/search", pc.Search)
	}
	tagPrivateGroup := privateGroup.Group("/tag")
	{
		tagPrivateGroup.POST("/create", pc.Create)
		tagPrivateGroup.PUT("/update/:id", pc.Update)
		tagPrivateGroup.DELETE("/:id", pc.Delete)
	}
}
