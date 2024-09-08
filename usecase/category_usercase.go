package usecase

import (
	"github.com/hxzhouh/go-zen.git/domain"
	"time"
)

type CategoryUsecase struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}

func (c CategoryUsecase) GetByCategoryID(id string) (domain.Category, error) {
	return c.GetByCategoryID(id)
}

func (c CategoryUsecase) Create(category *domain.Category) error {
	return c.categoryRepository.Create(category)
}

func (c CategoryUsecase) Update(category *domain.Category) error {
	return c.categoryRepository.Update(category)
}

func (c CategoryUsecase) Delete(id string) error {
	return c.categoryRepository.Delete(id)
}

func (c CategoryUsecase) GetAll() ([]domain.Category, error) {
	return c.categoryRepository.GetAll()
}

func (c CategoryUsecase) Search(keyword string) ([]domain.Category, error) {
	return c.categoryRepository.Search(keyword)
}

func NewCategoryUsecase(tagRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryRepository {
	return &CategoryUsecase{
		categoryRepository: tagRepository,
		contextTimeout:     timeout,
	}
}
