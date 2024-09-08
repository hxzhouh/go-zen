package sqlite

import (
	"context"
	"github.com/hxzhouh/go-zen.git/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		database: db,
	}
}
func (u UserRepository) Create(c context.Context, user *domain.User) error {
	return u.database.Create(user).Error
}

func (u UserRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	err := u.database.Find(&users).Error
	return users, err
}

func (u UserRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := u.database.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u UserRepository) GetByName(c context.Context, name string) (domain.User, error) {
	var user domain.User
	err := u.database.Where("name = ?", name).First(&user).Error
	return user, err
}

func (u UserRepository) GetByID(c context.Context, id string) (domain.User, error) {
	var user domain.User
	err := u.database.Where("id = ?", id).First(&user).Error
	return user, err
}
func (u UserRepository) DeleteByID(c context.Context, Id string) error {
	return u.database.Where("id = ?", Id).Delete(&domain.User{}).Error
}
