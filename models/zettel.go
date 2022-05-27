package models

import "gorm.io/gorm"

type Zettel struct {
	gorm.Model
	Title     string `gorm:"unique"`
	FilePath  string `gorm:"unique"`
	IsPrivate bool   `gorm:"default:false"`
}
