package main

import (
	"fmt"

	"github.com/ZDGharst/zet/models"
)

func Edit(zettels []models.Zettel) {
	if len(zettels) > 1 {
		fmt.Println(fmt.Sprintf("Found %d zettels matching that title", len(zettels)))
		for _, v := range zettels {
			fmt.Println("  " + v.Title)
		}
		return
	}
	if len(zettels) == 0 {
		fmt.Println("Found no zettels matching that title")
	}

	fmt.Println(zettels[0].Title)

	return
}
