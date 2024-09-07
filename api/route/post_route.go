package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/api/controller"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	"github.com/hxzhouh/go-zen.git/storage"
	"github.com/hxzhouh/go-zen.git/storage/sqlite"
	"github.com/hxzhouh/go-zen.git/usecase"
	"time"
)

func NewPostRouter(env *bootstrap.Env, timeout time.Duration, publicGroup *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	pr := sqlite.NewPostRepository(storage.DefaultStorage)
	pc := &controller.PostController{
		PostUsecase: usecase.NewPostUsecase(pr, timeout),
		Env:         nil,
	}
	postPublicGroup := publicGroup.Group("/posts")
	{
		postPublicGroup.GET("/list", pc.List)
		postPublicGroup.GET("/:id", pc.GetPostById)
	}
	postPrivateGroup := privateGroup.Group("/posts")
	{
		postPrivateGroup.POST("/create", pc.Create)
		postPrivateGroup.PUT("/update/:id", pc.Update)
		postPrivateGroup.POST("/upload", pc.Upload)
	}
}
