package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/zachgharst/zet/pkg/models"
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

	file, err := os.Create(readme)
	if err != nil {
		return err
	}

	if _, err := file.Write([]byte("# " + title)); err != nil {
		return err
	}

	file.Close()

	// Open README in vim
	cmd := exec.Command("vim", readme)
	cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	// Get actual title from file
	completedFile, err := os.Open(readme)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(completedFile)
	scanner.Scan()
	title = scanner.Text()[2:]
	completedFile.Close()

	// Git operations
	if err := Git_Sync(zpath, title); err != nil {
		return err
	}

	// Create row in database
	zettel := models.Zettel{Title: title, FilePath: now.Format("20060102150405")}
	if result := db.Create(&zettel); result.Error != nil {
		return result.Error
	}

	return nil
}
