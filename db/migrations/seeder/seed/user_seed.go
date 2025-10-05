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

func UserSeed(db *gorm.DB) error {
	log.Println("user seeder started")

	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	jsonFile, err := os.Open("./db/migrations/data/user.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)
	var users []entity.User
	json.Unmarshal(jsonData, &users)

	hasTable = db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	for _, user := range users {
		if err := db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&user).Error; err != nil {
			return err
		}
	}

	log.Println("user seeder finished")
	return nil
}
