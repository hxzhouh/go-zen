package usecase

import (
	"github.com/hxzhouh/go-zen.git/domain"
	"time"
)

type tagUsecase struct {
	tagtRepository domain.TagRepository
	contextTimeout time.Duration
}

func NewTagUsecase(tagRepository domain.TagRepository, timeout time.Duration) domain.TagUsecase {
	return &tagUsecase{
		tagtRepository: tagRepository,
		contextTimeout: timeout,
	}
}

func (t tagUsecase) SearchTag(name string) ([]domain.Tag, error) {
	return t.tagtRepository.SearchTag(name)
}

func (t tagUsecase) DeleteTag(id string) error {
	return t.tagtRepository.Delete(id)
}

func (t tagUsecase) UpdateTag(tag *domain.Tag) error {
	return t.tagtRepository.Update(tag)
}

func (t tagUsecase) CreateTag(tag *domain.Tag) error {
	return t.tagtRepository.Create(tag)
}

func (t tagUsecase) List() ([]domain.Tag, error) {
	return t.tagtRepository.GetAll()
}

func (t tagUsecase) GetByTagID(id string) (domain.Tag, error) {
	return t.tagtRepository.GetByID(id)
}

func (t tagUsecase) GetByTagIds(ids []string) ([]domain.Tag, error) {
	return t.tagtRepository.GetByIds(ids)
}
