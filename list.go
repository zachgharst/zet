package main

import (
	"fmt"

	"github.com/ZDGharst/zet/models"
	"gorm.io/gorm"
)

func List(db *gorm.DB) error {
	var zettels []models.Zettel
	result := db.Order("created_at desc").Find(&zettels)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Found", result.RowsAffected, "zettels")
	for _, zettel := range zettels {
		fmt.Println(" ", zettel.Title)
	}

	return nil
}
