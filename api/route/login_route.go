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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	ur := sqlite.NewUserRepository(storage.DefaultStorage)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
