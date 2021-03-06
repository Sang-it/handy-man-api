package gorm

import (
	"github.com/Sang-it/handy-man-api/internal/core/handy-man"
)

func (d *Database) CreateHandyManInDB(name string) {
	db := GetDB()
	handyman := handyman.HandyMan{FirstName: name}
	db.Create(&handyman)
}
