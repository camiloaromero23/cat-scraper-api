package types

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Name        string
	Link        string
	Image       string
	Description string
}
