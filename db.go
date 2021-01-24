package main

import (
	"github.com/ck46/ebus/store"
	"github.com/ck46/ebus/user"

	"github.com/jinzhu/gorm"
)

func dbMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&store.Category{},
		&store.Image{},
		&store.Department{},
		&store.Store{},
		&store.Item{},
	).Error
}
