package app

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/ZDGharst/zet/pkg/models"
	"gorm.io/gorm"
)

// Populate DB from the zettels directory by year
func Populate_DB(db *gorm.DB, zettels_directory, year string) error {
	var err error
	dirs, err := ioutil.ReadDir(zettels_directory + "/" + year)
	if err != nil {
		return err
	}

	var zettels []models.Zettel

	for _, dir := range dirs {
		file, err := os.Open(zettels_directory + "/" + year + "/" + dir.Name() + "/README.md")
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		var title = scanner.Text()[2:]
		createdAt, _ := time.Parse("20060102150405", dir.Name())

		zettels = append(zettels, models.Zettel{
			Model: gorm.Model{
				CreatedAt: createdAt,
				UpdatedAt: createdAt,
			},
			Title:     title,
			FilePath:  dir.Name(),
			IsPrivate: false,
		})
	}

	result := db.Create(&zettels)
	fmt.Println("Populated DB with", result.RowsAffected, "rows")

	return err
}
