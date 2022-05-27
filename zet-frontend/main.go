package main

import (
	"fmt"
	"os"

	"github.com/ZDGharst/zet/zet"
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
		// what if ZETTELS not set?
		zet.Create(os.Getenv("ZETTELS"), os.Args[2])
		return
	}

	fmt.Println("zet: command not found")
	fmt.Println("Try 'zet --help' for more information.")
}
