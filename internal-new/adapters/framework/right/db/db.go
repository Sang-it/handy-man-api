package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	DB *gorm.DB
}

func NewAdapter(connection_string string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB Failed to connect : %v", err)
	}

	return &Adapter{DB: db}, nil
}

func (dba Adapter) Migrate(model ...interface{}) {
	dba.DB.AutoMigrate(model)
}

func (dba Adapter) GetDB() *gorm.DB {
	return dba.DB
}
