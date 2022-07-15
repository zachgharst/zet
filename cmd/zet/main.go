package main

import (
	"fmt"
	"os"

	"github.com/ZDGharst/zet/pkg/app"
	"github.com/ZDGharst/zet/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	zettels_directory := os.Getenv("ZETTELS")
	if zettels_directory == "" {
		panic("ZETTELS env var not set")
	}

	db, err := gorm.Open(sqlite.Open(zettels_directory+"/.zetteldb"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Zettel{})

	if os.Args[1] == "create" {
		if err := app.Create(db, zettels_directory, os.Args[2]); err != nil {
			fmt.Println(err)
		}
		return
	}

	if os.Args[1] == "edit" {
		if err := app.Edit(db, zettels_directory, os.Args[2]); err != nil {
			fmt.Println(err)
		}
		return
	}

	if os.Args[1] == "list" {
		verbose := len(os.Args) > 2 && os.Args[2] == "-v"
		if err := app.ListAll(db, verbose); err != nil {
			fmt.Println(err)
		}
		return
	}

	if os.Args[1] == "populatedb" {
		// todo: hardcoded 2022
		if err := app.Populate_DB(db, zettels_directory, "2022"); err != nil {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("zet: command not found")
	fmt.Println("Try 'zet --help' for more information.")
}
