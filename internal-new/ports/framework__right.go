package ports

import "gorm.io/gorm"

type DatabasePort interface {
	GetDB() *gorm.DB
	Migrate(model ...interface{})
}
