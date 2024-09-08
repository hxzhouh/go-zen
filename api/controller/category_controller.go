package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	"github.com/hxzhouh/go-zen.git/domain"
)

type CategoryController struct {
	CategoryUsecase domain.CategoryRepository
	Env             *bootstrap.Env
}

func (c CategoryController) List(context *gin.Context) {

}

func (c CategoryController) Search(context *gin.Context) {

}

func (c CategoryController) Create(context *gin.Context) {

}

func (c CategoryController) Update(context *gin.Context) {

}

func (c CategoryController) Delete(context *gin.Context) {

}
