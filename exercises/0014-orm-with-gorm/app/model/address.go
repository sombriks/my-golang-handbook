package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street    string
	ContactID uint
}
