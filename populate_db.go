package main

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/ZDGharst/zet/models"
	"gorm.io/gorm"
)

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
		var firstln = scanner.Text()[2:]

		zettels = append(zettels, models.Zettel{
			Title:    firstln,
			FilePath: dir.Name(),
		})
	}

	db.Create(&zettels)

	return err
}
