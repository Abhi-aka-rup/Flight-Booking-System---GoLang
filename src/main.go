package main

import (
	"gorm.io/gorm"

	"flight-system/src/databaseInit"
)

var db *gorm.DB

func main() {
	db = databaseInit.InitDb()
}
