package seed

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PostSeed(db *gorm.DB) error {
	log.Println("post seeder started")

	hasTable := db.Migrator().HasTable(&entity.Post{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Post{}); err != nil {
			return err
		}
	}

	jsonFile, err := os.Open("./db/migrations/data/post.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)
	var posts []entity.Post
	json.Unmarshal(jsonData, &posts)

	hasTable = db.Migrator().HasTable(&entity.Post{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Post{}); err != nil {
			return err
		}
	}

	for _, post := range posts {
		if err := db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&post).Error; err != nil {
			return err
		}
	}

	log.Println("post seeder finished")
	return nil
}
