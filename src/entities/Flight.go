package entities

import "gorm.io/gorm"

type Flight struct {
	gorm.Model

	Origin      string
	Destination string
	Price       float32
}
