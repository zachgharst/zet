package main

import (
	"fmt"
	"os"

	"github.com/ZDGharst/zet/models"
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
		if err := Create(db, zettels_directory, os.Args[2]); err != nil {
			fmt.Println(err)
		}
		return
	}

	if os.Args[1] == "edit" {
		zettels, err := FindByTitle(db, os.Args[2])
		if err != nil {
			fmt.Println(err)
		}
		Edit(zettels)

		return
	}

	if os.Args[1] == "list" {
		if err := List(db); err != nil {
			fmt.Println(err)
		}
		return
	}

	if os.Args[1] == "populatedb" {
		// todo: hardcoded 2022
		if err := Populate_DB(db, zettels_directory, "2022"); err != nil {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("zet: command not found")
	fmt.Println("Try 'zet --help' for more information.")
}
