package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ZDGharst/zet/zet/models"
)

func List(zettels_directory string) error {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	var zettels []models.Zettel
	result := db.Find(&zettels)
	if result.Error != nil {
		return result.Error
	}

	for _, zettel := range zettels {
		fmt.Println(zettel.Title)
	}

	return nil
}
