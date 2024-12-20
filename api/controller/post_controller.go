package controller

import (
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hxzhouh/go-zen.git/bootstrap"
	"github.com/hxzhouh/go-zen.git/domain"
	"github.com/hxzhouh/go-zen.git/internal"
)

type PostController struct {
	PostUsecase domain.PostUsecase
	Env         *bootstrap.Env
}

// List 分页获取所有文章
// @Summary 分页获取所有文章
// @Description 根据id获取文章，如果为空返回错误
// @Param request_id header string true "Request ID"
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Tags 文章相关接口
// @Accept application/json
// @Produce application/json
// @Param id path string true "文章id"
// @Success 200 {object} []domain.Post
// @Router /posts/:id [get]
func (pc *PostController) List(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	// 转换为整数类型
	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	prevOffset := offset - limit
	if prevOffset < 0 {
		prevOffset = 0 // 确保不为负数
	}
	nextOffset := offset + limit
	posts, err := pc.PostUsecase.List(offset, limit)
	if err != nil {
		slog.Error("List error", "error", err.Error())
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	// 判断是否有上一页
	hasPrev := offset > 0

	// 判断是否有下一页（需要知道总的文章数 totalCount）
	hasNext := offset+limit < 123
	c.HTML(http.StatusOK, "posts_list.html", gin.H{
		"Posts":      posts,
		"Limit":      limit,
		"Offset":     offset,
		"PrevOffset": prevOffset,
		"NextOffset": nextOffset,
		"HasPrev":    hasPrev,
		"HasNext":    hasNext,
	})
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
	post.ContentHtml = internal.MdToHTML([]byte(post.Content))
	c.HTML(http.StatusOK, "post_detail.html", gin.H{
		"Post": post,
	})
}

// Create 创建文章
// @Summary 创建文章
// @Description 创建文章，返回文章Id
// @Tags 文章相关接口
// @Param   request_id  header  string  true  "Request ID"
// @Param   Authorization  header  string  true  "JWT"
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
	userID := context.GetString("x-user-id")
	postId, err := pc.PostUsecase.CreatePost(userID, &postRequest)
	if err != nil {
		slog.Error("PostController Create error", "error", err.Error())
		context.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	context.JSON(http.StatusOK, &domain.CreatePostResponse{ID: postId})
}

// Upload 从文件上传创建文章
// @Summary 从文件上传创建文章
// @Description 从文件上传创建文章，返回文章Id
// @Tags 文章相关接口
// @Param   request_id  header  string  true  "Request ID"
// @Param   Authorization  header  string  true  "JWT"
// @Accept multipart/form-data
// @Param mdFile formData file true "需要上传的markdown文件"
// @Produce application/json
// @Success 200 {object} domain.CreatePostResponse
// @Router /posts/upload [post]
func (pc *PostController) Upload(context *gin.Context) {
	file, _, err := context.Request.FormFile("mdFile")
	if err != nil {
		slog.Error("PostController Upload error", "error", err.Error())
		context.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	////获取文件名
	//filename := header.Filename
	defer func() {
		_ = file.Close()
	}()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		slog.Error("PostController Upload error", "error", err.Error())
		context.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	postRequest := domain.CreatePostRequest{
		Title:    "",
		SubTitle: "",
		Summary:  "",
		Cover:    "",
		Content:  string(fileBytes),
		//Tags:     nil,
	}
	postId, err := pc.PostUsecase.CreatePost("", &postRequest)
	if err != nil {
		slog.Error("PostController Upload error", "error", err.Error())
		context.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	context.JSON(http.StatusOK, &domain.CreatePostResponse{ID: postId})
}

func (pc *PostController) Update(context *gin.Context) {

}
