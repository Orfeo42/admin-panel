package data

import "gorm.io/gorm"

type CustomerModel struct {
	gorm.Model
	Name    string
	Surname string
}
