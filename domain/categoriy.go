package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryId string `gorm:"type:varchar(32)"`
	Name       string `gorm:"type:varchar(32)"`
	Summary    string `gorm:"type:varchar(255)"`
}

type CategoryRepository interface {
	Create(category *Category) error
	Update(category *Category) error
	Delete(id string) error
	GetAll() ([]Category, error)
	Search(keyword string) ([]Category, error)
	GetByCategoryID(id string) (Category, error)
}
