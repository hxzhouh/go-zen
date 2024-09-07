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
	err := DefaultStorage.AutoMigrate(&domain.User{})
	if err != nil {
		return err
	}
	err = DefaultStorage.AutoMigrate(&domain.Post{})
	if err != nil {
		return err
	}
	return nil
}
