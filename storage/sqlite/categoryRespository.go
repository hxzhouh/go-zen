package sqlite

import (
	"github.com/hxzhouh/go-zen.git/domain"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	database *gorm.DB
}

func (u CategoryRepository) GetByCategoryID(id string) (domain.Category, error) {
	c := domain.Category{}
	err := u.database.Where("category_id = ?", id).First(&c).Error
	return c, err
}

func (u CategoryRepository) Search(keyword string) ([]domain.Category, error) {
	var c []domain.Category
	err := u.database.Where("name LIKE ?", "%"+keyword+"%").Or("summary LIKE ?", "%"+keyword+"%").Find(&c).Error
	return c, err
}

func (u CategoryRepository) Create(category *domain.Category) error {
	return u.database.Create(category).Error
}

func (u CategoryRepository) Update(category *domain.Category) error {
	return u.database.Save(category).Error
}

func (u CategoryRepository) GetAll() ([]domain.Category, error) {
	var c []domain.Category
	err := u.database.Find(&c).Error
	return c, err
}

func (u CategoryRepository) GetByID(id string) (domain.Category, error) {
	c := domain.Category{}
	err := u.database.Where("category_id = ?", id).First(&c).Error
	return c, err
}

func (u CategoryRepository) GetByIds(ids []string) ([]domain.Category, error) {
	var category []domain.Category
	err := u.database.Where("category_id IN (?)", ids).Find(&category).Error
	return category, err
}

func (u CategoryRepository) Delete(id string) error {
	return u.database.Where("category_id = ?", id).Delete(&domain.Post{}).Error
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return CategoryRepository{
		database: db,
	}
}
