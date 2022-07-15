package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"gorm.io/gorm"
)

func Edit(db *gorm.DB, zettels_directory string, title string) error {
	zettels, err := FindByTitle(db, title)
	if err != nil {
		return err
	}

	if len(zettels) > 1 {
		fmt.Println(fmt.Sprintf("Found %d zettels matching \"%s\"", len(zettels), title))
		for _, v := range zettels {
			fmt.Println("  " + v.Title)
		}
		return errors.New("Please refine your search.")
	}

	if len(zettels) == 0 {
		return errors.New("Found no zettels matching that title.")
	}

	zettel := zettels[0]

	// Path to README
	zpath := fmt.Sprintf(
		"%s/%s/%s",
		zettels_directory,
		zettel.FilePath[:4],
		zettel.FilePath,
	)
	readme := fmt.Sprintf("%s/README.md", zpath)

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

	// Edit row in database
	if zettel.Title != title {
		zettel.Title = title
		if result := db.Save(&zettel); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
