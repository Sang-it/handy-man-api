package database

import (
	"github.com/Sang-it/handy-man-api/ports/gorm"
)

func GetDatabaseAdapter() gorm.DatabaseActions {
	return gorm.GetDatabaseAdapter()
}
