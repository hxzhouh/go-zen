package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/api/controller"
	"github.com/hxzhouh/go-zen.git/storage"
	"github.com/hxzhouh/go-zen.git/storage/sqlite"
	"github.com/hxzhouh/go-zen.git/usecase"
	"time"
)

func NewCategoryRouter(timeout time.Duration, publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	pr := sqlite.NewCategoryRepository(storage.DefaultStorage)
	pc := &controller.CategoryController{
		CategoryUsecase: usecase.NewCategoryUsecase(pr, timeout),
		Env:             nil,
	}
	tagPublicGroup := publicGroup.Group("/category")
	{
		tagPublicGroup.GET("/list", pc.List)
		tagPublicGroup.GET("/search", pc.Search)
	}
	tagPrivateGroup := privateGroup.Group("/category")
	{
		tagPrivateGroup.POST("/create", pc.Create)
		tagPrivateGroup.PUT("/update/:id", pc.Update)
		tagPrivateGroup.DELETE("/delete/:id", pc.Delete)
	}
}
