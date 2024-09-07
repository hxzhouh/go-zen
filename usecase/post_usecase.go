package usecase

import (
	"crypto/md5"
	"fmt"
	"github.com/hxzhouh/go-zen.git/domain"
	"github.com/hxzhouh/go-zen.git/internal/markdown"
	"time"
)

type postUsecase struct {
	postRepository domain.PostRepository
	contextTimeout time.Duration
}

func NewPostUsecase(postRepository domain.PostRepository, timeout time.Duration) domain.PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
		contextTimeout: timeout,
	}
}

func (p postUsecase) List(offset, limit int) ([]domain.Post, error) {
	return p.postRepository.Fetch(offset, limit)
}

func (p postUsecase) GetByID(id string) (domain.Post, error) {
	return p.postRepository.GetByID(id)
}

func (p postUsecase) SearchByKeyword(keyword string, offset, limit int) ([]domain.Post, error) {
	return p.postRepository.Search(keyword, offset, limit)
}

func (p postUsecase) CreatePost(authorID string, postReq *domain.CreatePostRequest) error {
	post := &domain.Post{
		Title:       postReq.Title,
		SubTitle:    postReq.SubTitle,
		Summary:     postReq.Summary,
		Cover:       postReq.Cover,
		Content:     postReq.Content,
		ContentHtml: string(markdown.MdToHTML([]byte(postReq.Content))),
		Tags:        postReq.Tags,
		Md5:         calcPostMd5(postReq),
		AuthorID:    authorID,
	}
	return p.postRepository.Create(post)
}

func calcPostMd5(post *domain.CreatePostRequest) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(post.Title+post.Content+post.SubTitle)))
}
