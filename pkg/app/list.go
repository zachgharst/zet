package app

import (
	"fmt"

	"github.com/zachgharst/zet/pkg/models"
	"gorm.io/gorm"
)

func ListAll(db *gorm.DB, verbose bool) error {
	var zettels []models.Zettel
	result := db.Order("created_at desc").Find(&zettels)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Found", result.RowsAffected, "zettels")
	for index, zettel := range zettels {
		if verbose {
			fmt.Printf(
				"  #%d: %s - %s\n",
				len(zettels)-index,
				zettel.Title,
				zettel.FilePath,
			)
		} else {
			fmt.Println("  ", zettel.Title)
		}
	}

	return nil
}
