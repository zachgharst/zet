package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/ZDGharst/zet/models"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, zettels_directory, title string) error {
	now := time.Now()

	// Check if the year folder is already created, create if not
	zpath := fmt.Sprintf("%s/%d", zettels_directory, now.Year())
	if _, err := os.Stat(zpath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(zpath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create the zettel folder and file.
	zpath = fmt.Sprintf("%s/%s", zpath, now.Format("20060102150405"))
	readme := zpath + "/README.md"
	if err := os.Mkdir(zpath, os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Create(readme); err != nil {
		return err
	}

	if err := os.WriteFile(readme, []byte("# "+title), 0644); err != nil {
		return err
	}

	// Open README in vim
	cmd := exec.Command("vim", readme)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	// Git operations
	if err := Git_Sync(zpath); err != nil {
		return err
	}

	// Create row in database
	zettel := models.Zettel{Title: title, FilePath: zpath}
	if result := db.Create(&zettel); result.Error != nil {
		return result.Error
	}

	return nil
}
