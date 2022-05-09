package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func help() {
	//todo: implement help
	fmt.Println("todo: implement help")
}

func create(title string) {
	//todo: what if env not set?
	zenv := os.Getenv("ZETTELS")
	now := time.Now()

	// Check if the year folder is already created, create if not
	zpath := fmt.Sprintf("%s/%d", zenv, now.Year())
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

	gc := exec.Command("git", "commit", "-m", title)
	if err := gc.Run(); err != nil {
		fmt.Println(err)
		return
	}

	gpush := exec.Command("git", "push")
	if err := gpush.Run(); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	if os.Args[1] == "create" {
		create(os.Args[2])
		return
	}

	fmt.Println("zet: command not found")
	fmt.Println("Try 'zet --help' for more information.")
}
