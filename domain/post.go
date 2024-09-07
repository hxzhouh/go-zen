package domain

type CreatePostRequest struct {
	Title    string   `json:"title" binding:"required"`
	SubTitle string   `json:"subTitle"`
	Summary  string   `json:"summary"`
	Cover    string   `json:"cover"`
	Content  string   `json:"content" binding:"required"`
	Tags     []Tag    `json:"tag"`
	Category Category `json:"category"`
}
type CreatePostResponse struct {
	ID string `json:"id"`
}

type Post struct {
	ID          string   `gorm:"primaryKey;autoIncrement"`
	Title       string   `gorm:"type:varchar(255)"`
	SubTitle    string   `gorm:"type:varchar(255)"`
	Summary     string   `gorm:"type:varchar(255)"`
	Draft       bool     `gorm:"type:boolean"`
	Cover       string   `gorm:"type:varchar(255)"`
	Content     string   `gorm:"type:text"`
	ContentHtml string   `gorm:"type:text"`
	AuthorID    string   `gorm:"type:varchar(255)"`
	Md5         string   `gorm:"type:varchar(64)"`
	Tags        []Tag    `gorm:"many2many:post_tags;"`
	Category    Category `gorm:"foreignKey:ID;references:CategoryID"`
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
	CreatePost(authorID string, post *CreatePostRequest) error
	List(offset, limit int) ([]Post, error)
	GetByID(id string) (Post, error)
	SearchByKeyword(keyword string, offset, limit int) ([]Post, error)
}
