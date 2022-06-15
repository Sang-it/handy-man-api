package handyman

import (
	"gorm.io/gorm"
)

type HandyMan struct {
	gorm.Model
	FirstName     string
	MiddleName    string
	LastName      string
	ContactNumber string
	Email         string
	Body          string
}
