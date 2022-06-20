package database

import (
	"fmt"
	"github.com/Sang-it/handy-man-api/configs/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var (
	DB *gorm.DB
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(environment.Get("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
