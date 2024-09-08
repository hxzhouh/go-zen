package sqlite

import (
	"github.com/hxzhouh/go-zen.git/domain"
	"gorm.io/gorm"
)

type TagRepository struct {
	database *gorm.DB
}

func (u TagRepository) Create(post *domain.Tag) error {
	return u.database.Create(post).Error
}

func (u TagRepository) Update(tag *domain.Tag) error {
	return u.database.Save(tag).Error
}

func (u TagRepository) GetAll() ([]domain.Tag, error) {
	tags := []domain.Tag{}
	err := u.database.Find(&tags).Error
	return tags, err
}

func (u TagRepository) GetByID(id string) (domain.Tag, error) {
	tag := domain.Tag{}
	err := u.database.Where("tag_id = ?", id).First(&tag).Error
	return tag, err
}

func (u TagRepository) GetByIds(ids []string) ([]domain.Tag, error) {
	tags := []domain.Tag{}
	err := u.database.Where("tag_id IN (?)", ids).Find(&tags).Error
	return tags, err
}

func (u TagRepository) Delete(id string) error {
	return u.database.Where("tag_id = ?", id).Delete(&domain.Tag{}).Error
}

func (u TagRepository) SearchTag(name string) ([]domain.Tag, error) {
	var tags []domain.Tag
	err := u.database.Where("name LIKE ?", "%"+name+"%").Find(&tags).Error
	return tags, err
}

func NewTagRepository(db *gorm.DB) domain.TagRepository {
	return TagRepository{
		database: db,
	}
}
