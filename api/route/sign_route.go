package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/api/controller"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	"github.com/hxzhouh/go-zen.git/storage/sqlite"
	"github.com/hxzhouh/go-zen.git/usecase"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := sqlite.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)

	group.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"content": "You reached the go-zen blog site...",
		})
	})
}
