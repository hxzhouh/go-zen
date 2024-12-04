package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/api/middleware"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	_ "github.com/hxzhouh/go-zen.git/docs"
	"github.com/hxzhouh/go-zen.git/storage"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")
	publicRouter.Use(middleware.RequestLogMiddleware())
	privateRouter := gin.Group("")
	privateRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret), middleware.RequestLogMiddleware())

	NewSignupRouter(env, timeout, storage.DefaultStorage, publicRouter)
	NewLoginRouter(env, timeout, publicRouter)
	NewIndexRoute(publicRouter)
	NewPostRouter(env, timeout, publicRouter, privateRouter)
	NewTagRouter(timeout, publicRouter, privateRouter)

	if env.AppEnv == "development" {
		NewSwaggerRouter(publicRouter)
	}
}
