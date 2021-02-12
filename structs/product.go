package structs

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    string
	Name  string
	Price int
}
