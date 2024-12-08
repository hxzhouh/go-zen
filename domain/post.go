package domain

import (
	"database/sql/driver"
	"encoding/json"
	"html/template"

	"gorm.io/gorm"
)

type CreatePostRequest struct {
	Title      string   `json:"title" binding:"required"`
	SubTitle   string   `json:"subTitle"`
	Summary    string   `json:"summary"`
	Cover      string   `json:"cover"`
	Content    string   `json:"content" binding:"required"`
	TagIds     []string `json:"tag_ids"`
	CategoryId []string `json:"category_ids"`
}
type CreatePostResponse struct {
	ID string `json:"id"`
}
type TagList []string
type CategoryList []string

func (p TagList) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现方法
func (p *TagList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}

func (p CategoryList) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan 实现方法
func (p *CategoryList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}

type Post struct {
	gorm.Model
	PostId      string        `gorm:"unique"`
	Title       string        `gorm:"type:varchar(255)"`
	SubTitle    string        `gorm:"type:varchar(255)"`
	Summary     string        `gorm:"type:varchar(255)"`
	Draft       bool          `gorm:"type:boolean"`
	Cover       string        `gorm:"type:varchar(255)"`
	Content     string        `gorm:"type:text"`
	ContentHtml template.HTML `gorm:"type:text"`
	AuthorID    string        `gorm:"type:varchar(255)"`
	Md5         string        `gorm:"type:varchar(64)"`
	TagIds      TagList       `json:"tag_ids"`
	CategoryId  CategoryList  `json:"category_id"`
	Reads       int           `gorm:"type:int"`
	Likes       int           `gorm:"type:int"`
}

type PostRepository interface {
	Create(post *Post) error
	Fetch(offset, limit int) ([]Post, error)
	GetByID(id string) (Post, error)
	Update(post *Post) error
	Delete(id string) error
	Search(keyword string, offset, limit int) ([]Post, error)
	GetByTag(tag string) ([]Post, error)
	GetByCategory(category string) ([]Post, error)
}

type PostUsecase interface {
	CreatePost(authorID string, post *CreatePostRequest) (string, error)
	List(offset, limit int) ([]Post, error)
	GetByID(id string) (Post, error)
	SearchByKeyword(keyword string, offset, limit int) ([]Post, error)
}
