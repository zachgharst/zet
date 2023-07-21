package app

import (
	"github.com/zachgharst/zet/pkg/models"
	"gorm.io/gorm"
)

// Populate DB from the zettels directory by year
func FindByTitle(db *gorm.DB, title string) ([]models.Zettel, error) {
	var err error
	var zettels []models.Zettel
	db.Where("title LIKE ?", "%"+title+"%").Find(&zettels)
	return zettels, err
}
