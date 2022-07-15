package main

import (
	"fmt"

	"github.com/ZDGharst/zet/models"
	"gorm.io/gorm"
)

func List(db *gorm.DB, verbose bool) error {
	var zettels []models.Zettel
	result := db.Order("created_at desc").Find(&zettels)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Found", result.RowsAffected, "zettels")
	for index, zettel := range zettels {
		if verbose {
			lineStr := fmt.Sprintf(
				"  #%d: %s - %s",
				len(zettels)-index,
				zettel.Title,
				zettel.FilePath,
			)
			fmt.Println(lineStr)
		} else {
			fmt.Println("  ", zettel.Title)
		}
	}

	return nil
}
