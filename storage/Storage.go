package storage

import (
	"errors"
	"github.com/hxzhouh/go-zen.git/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DefaultStorage *gorm.DB

func InitStorage(dataType string, dsn string) error {
	//var db *gorm.DB
	var err error
	switch dataType {
	case "sqlite3":
		DefaultStorage, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		err = migration()
		return nil
	}
	return errors.New("unsupported data type")
}

func migration() error {
	return DefaultStorage.AutoMigrate(&domain.User{}, &domain.Post{}, &domain.Tag{}, &domain.Category{})
}
