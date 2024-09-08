package domain

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagId string `gorm:"unique"`
	Name  string `gorm:"type:varchar(255)"`
}

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}
type ListTagsResponse struct {
	Tags []Tag `json:"tags"`
}
type TagRepository interface {
	Create(tag *Tag) error
	Update(tag *Tag) error
	Delete(id string) error
	GetAll() ([]Tag, error)
	GetByID(id string) (Tag, error)
	GetByIds(ids []string) ([]Tag, error)
	SearchTag(name string) ([]Tag, error)
}

type TagUsecase interface {
	CreateTag(tag *Tag) error
	List() ([]Tag, error)
	GetByTagID(id string) (Tag, error)
	GetByTagIds(ids []string) ([]Tag, error)
	UpdateTag(tag *Tag) error
	SearchTag(name string) ([]Tag, error)
	DeleteTag(id string) error
}
