package sqlite

import (
	"github.com/hxzhouh/go-zen.git/domain"
	"gorm.io/gorm"
)

type PostRepository struct {
	database *gorm.DB
}

func (u PostRepository) Delete(id string) error {
	return u.database.Where("id = ?", id).Delete(&domain.Post{}).Error
}

func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return PostRepository{
		database: db,
	}
}

func (u PostRepository) Create(post *domain.Post) error {
	return u.database.Create(post).Error
}

func (u PostRepository) Fetch(offset, limit int) ([]domain.Post, error) {
	var posts []domain.Post
	err := u.database.Limit(limit).Offset(offset).Find(&posts).Error
	return posts, err
}

func (u PostRepository) GetByID(id string) (domain.Post, error) {
	var post domain.Post
	err := u.database.Where("post_id = ?", id).First(&post).Error
	return post, err
}

func (u PostRepository) Update(post *domain.Post) error {
	return u.database.Save(post).Error
}

func (u PostRepository) DeleteByID(id string) error {
	return u.database.Where("post_id = ?", id).Delete(&domain.Post{}).Error

}

func (u PostRepository) Search(keyword string, offset, limit int) ([]domain.Post, error) {
	var posts []domain.Post
	err := u.database.Where("title LIKE ?", "%"+keyword+"%").Or("content LIKE ?", "%"+keyword+"%").Find(&posts).Error
	return posts, err
}

func (u PostRepository) GetByTag(tag string) ([]domain.Post, error) {
	var posts []domain.Post
	err := u.database.Where("tag = ?", tag).Find(&posts).Error
	return posts, err
}

func (u PostRepository) GetByCategory(category string) ([]domain.Post, error) {
	var posts []domain.Post
	err := u.database.Where("category = ?", category).Find(&posts).Error
	return posts, err
}
