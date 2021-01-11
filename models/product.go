package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	name       string
	image      byte
	price      uint32
	categoryId uint
}
