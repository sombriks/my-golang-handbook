package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name      string
	Addresses []Address
}
