package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	"github.com/hxzhouh/go-zen.git/domain"
	"log/slog"
	"net/http"
)

type PostController struct {
	PostUsecase domain.PostUsecase
	Env         *bootstrap.Env
}

func (pc *PostController) List(c *gin.Context) {

}

// GetPostById 获取文章
// @Summary 根据id获取文章
// @Description 根据id获取文章，如果为空返回错误
// @Param request_id header string true "Request ID"
// @Tags 文章相关接口
// @Accept application/json
// @Produce application/json
// @Param id path string true "文章id"
// @Success 200 {object} domain.Post
// @Router /posts/:id [get]
func (pc *PostController) GetPostById(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		slog.Error("GetPostById id is empty")
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "id is empty"})
		return
	}
	post, err := pc.PostUsecase.GetByID(id)
	if err != nil {
		slog.Error("GetPostById error", "error", err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// Create 创建文章
// @Summary 创建文章
// @Description 创建文章，返回文章Id
// @Tags 文章相关接口
// @Param   request_id  header  string  true  "Request ID"
// @security Authorization
// @Accept application/json
// @Produce application/json
// @Param post body domain.CreatePostRequest true "创建文章的参数"
// @Success 200 {object} domain.CreatePostResponse
// @Router /posts/create [post]
func (pc *PostController) Create(context *gin.Context) {
	postRequest := domain.CreatePostRequest{}
	err := context.ShouldBind(&postRequest)
	if err != nil {
		slog.Error("PostController Create error", "error", err.Error())
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	// todo convert to domain.Post
	post := domain.Post{}
	err = pc.PostUsecase.CreatePost("", &postRequest)
	if err != nil {
		slog.Error("PostController Create error", "error", err.Error())
		context.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	context.JSON(http.StatusOK, &domain.CreatePostResponse{ID: post.ID})
}

// Upload 从文件上传创建文章
// @Summary 从文件上传创建文章
// @Description 从文件上传创建文章，返回文章Id
// @Tags 文章相关接口
// @Param   request_id  header  string  true  "Request ID"
// @security Authorization
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Produce application/json
// @Success 200 {object} domain.CreatePostResponse
// @Router /posts/upload [post]
func (pc *PostController) Upload(context *gin.Context) {
	
}
func (pc *PostController) Update(context *gin.Context) {

}
