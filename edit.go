package main

import (
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

	readme := fmt.Sprintf(
		"%s/%s/%s/README.md",
		zettels_directory,
		zettel.FilePath[:4],
		zettel.FilePath,
	)

	cmd := exec.Command("vim", readme)
	cmd.Stdin, cmd.Stdout = os.Stdin, os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
