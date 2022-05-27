package main

import (
	"fmt"
	"os"
)

func help() {
	//todo: implement help
	fmt.Println("todo: implement help")
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	if os.Args[1] == "create" {
		// todo: what if ZETTELS not set?
		if err := Create(os.Getenv("ZETTELS"), os.Args[2]); err != nil {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("zet: command not found")
	fmt.Println("Try 'zet --help' for more information.")
}
