package main

import (
	"bufio"
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
		fmt.Println(fmt.Sprintf("Found %d zettels matching that title", len(zettels)))
		for _, v := range zettels {
			fmt.Println("  " + v.Title)
		}
		return nil
	}

	if len(zettels) == 0 {
		fmt.Println("Found no zettels matching that title")
		return nil
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
	zettel.Title = title
	if result := db.Save(&zettel); result.Error != nil {
		return result.Error
	}

	return nil
}
