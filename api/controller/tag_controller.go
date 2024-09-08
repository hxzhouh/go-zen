package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	"github.com/hxzhouh/go-zen.git/domain"
	"github.com/hxzhouh/go-zen.git/utils"
)

type TagController struct {
	TagUsecase domain.TagUsecase
	Env        *bootstrap.Env
}

// Update 更新Tag
// @Summary 更新Tag
// @Description 创建Tag，返回Tag
// @Tags Tag相关接口
// @Param   request_id  header  string  true  "Request ID"
// @Param   Authorization  header  string  true  "JWT"
// @Accept application/json
// @Produce application/json
// @Param post body domain.Tag true "tag"
// @Success 200 {object} domain.Tag "Tag"
// @Router /tag/update [post]
func (c TagController) Update(context *gin.Context) {
	value := domain.Tag{}
	if err := context.ShouldBindJSON(&value); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := c.TagUsecase.UpdateTag(&value)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, value)
}

// Create 创建Tag
// @Summary 创建Tag
// @Description 创建Tag，返回Tag
// @Tags Tag相关接口
// @Param   request_id  header  string  true  "Request ID"
// @Param   Authorization  header  string  true  "JWT"
// @Accept application/json
// @Produce application/json
// @Param post body domain.CreateTagRequest true "tagName"
// @Success 200 {object} domain.Tag "Tag"
// @Router /tag/create [post]
func (c TagController) Create(context *gin.Context) {
	value := domain.CreateTagRequest{}
	if err := context.ShouldBindJSON(&value); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tag := domain.Tag{
		Name:  value.Name,
		TagId: utils.GenerateSnowflakeID().Base32(),
	}
	err := c.TagUsecase.CreateTag(&tag)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, tag)
}

// List 获取所有tag
// @Summary 获取所有tag
// @Description
// @Tags Tag相关接口
// @Param   request_id  header  string  true  "Request ID"
// @Accept application/json
// @Produce application/json
// @Success 200 {object} domain.ListTagsResponse "Tag"
// @Router /tag/list [get]
func (c TagController) List(context *gin.Context) {
	var tags []domain.Tag
	tags, err := c.TagUsecase.List()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp := domain.ListTagsResponse{
		Tags: tags,
	}
	context.JSON(200, resp)
}

// Search 根据关键字查询tag
// @Summary 根据关键字查询tag
// @Description
// @Tags Tag相关接口
// @Param   request_id  header  string  true  "Request ID"
// @Accept application/json
// @Produce application/json
// @Param keyword query string true "keyword"
// @Success 200 {object} domain.ListTagsResponse "Tags"
// @Router /tag/search [get]
func (c TagController) Search(context *gin.Context) {
	keyword := context.Query("keyword")
	tags, err := c.TagUsecase.SearchTag(keyword)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp := domain.ListTagsResponse{
		Tags: tags,
	}
	context.JSON(200, resp)
}

// Delete 删除tag
// @Summary 根据Id 删除tag
// @Description
// @Tags Tag相关接口
// @Param   request_id  header  string  true  "Request ID"
// @Param   Authorization  header  string  true  "JWT"
// @Accept application/json
// @Produce application/json
// @Param id query string true "id"
// @Router /tag/search [delete]
func (c TagController) Delete(context *gin.Context) {
	id := context.Param("id")
	err := c.TagUsecase.DeleteTag(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, gin.H{"message": "success"})
}
