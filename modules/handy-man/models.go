package handyman

import (
	"gorm.io/gorm"
)

type HandyMan struct {
	gorm.Model
	FirstName string
}
