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

	db, err := gorm.Open(sqlite.Open("zettel.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Zettel{})

	if os.Args[1] == "create" {
		// todo: what if ZETTELS not set?
		if err := Create(db, os.Getenv("ZETTELS"), os.Args[2]); err != nil {
			fmt.Println(err)
		}
		return
	}

	if os.Args[1] == "list" {
		// todo: what if ZETTELS not set?
		if err := List(db); err != nil {
			fmt.Println(err)
		}
		return
	}

	if os.Args[1] == "populatedb" {
		// todo: what if ZETTELS not set?
		if err := Populate_DB(db, os.Getenv("ZETTELS"), "2022"); err != nil {
			fmt.Println(err)
		}
		return
	}

	fmt.Println("zet: command not found")
	fmt.Println("Try 'zet --help' for more information.")
}
