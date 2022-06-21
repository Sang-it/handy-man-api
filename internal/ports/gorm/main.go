package gorm

import (
	"fmt"
	"github.com/Sang-it/handy-man-api/configs/environment"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

type DatabaseActions interface {
	CreateHandyManInDB(name string)
	Migrate(model ...interface{})
	Init() *gorm.DB
}

var (
	DB *gorm.DB
)

func (d *Database) Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(environment.Get("CONNECTION_STRING")), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	DB = db
	return DB
}

func (d *Database) Migrate(model ...interface{}) {
	DB.AutoMigrate(model)
}

func GetDB() *gorm.DB {
	return DB
}

func GetDatabaseAdapter() DatabaseActions {
	return &Database{}
}
