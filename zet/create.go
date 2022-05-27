package zet

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Create(zettels_directory, title string) {
	now := time.Now()

	// Check if the year folder is already created, create if not
	zpath := fmt.Sprintf("%s/%d", zettels_directory, now.Year())
	if _, err := os.Stat(zpath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(zpath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Create the zettel folder and file.
	zpath = fmt.Sprintf("%s/%s", zpath, now.Format("20060102150405"))
	readme := zpath + "/README.md"
	if err := os.Mkdir(zpath, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Create(readme); err != nil {
		fmt.Println(err)
		return
	}

	if err := os.WriteFile(readme, []byte("# "+title), 0644); err != nil {
		fmt.Println(err)
		return
	}

	// Open README in vim
	cmd := exec.Command("vim", readme)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// Git operations
	os.Chdir(zpath)

	gpull := exec.Command("git", "pull")
	if err := gpull.Run(); err != nil {
		fmt.Println(err)
		return
	}

	ga := exec.Command("git", "add", ".")
	if err := ga.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// gc := exec.Command("git", "commit", "-m", title)
	// if err := gc.Run(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// gpush := exec.Command("git", "push")
	// if err := gpush.Run(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}