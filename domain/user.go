package domain

import "context"

const (
	CollectionUser = "users"
)

type User struct {
	ID       string `gorm:"primaryKey;autoIncrement"`
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	DeleteByID(c context.Context, id string) error
}
