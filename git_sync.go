package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Git_Sync(directory string) error {
	if err := os.Chdir(directory); err != nil {
		return err
	}
	fmt.Println("cd", directory)

	gpull := exec.Command("git", "pull")
	gpull.Stdout = os.Stdout
	if err := gpull.Run(); err != nil {
		return err
	}

	ga := exec.Command("git", "add", ".")
	if err := ga.Run(); err != nil {
		return err
	}

	// gc := exec.Command("git", "commit", "-m", title)
	// gc.Stdout = os.Stdout
	// if err := gc.Run(); err != nil {
	// 	return err
	// }

	// gpush := exec.Command("git", "push")
	// gpush.Stdout = os.Stdout
	// if err := gpush.Run(); err != nil {
	// 	return err
	// }

	return nil
}
