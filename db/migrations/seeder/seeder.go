package seeder

import (
	"log"
	"os"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/db/migrations/seeder/seed"
	"gorm.io/gorm"
)

func RunSeeder(db *gorm.DB) error {
	log.Println("seeder started")

	app := os.Getenv("APP_ENV")
	if app != "production" {
		if err := seed.UserSeed(db); err != nil {
			log.Fatalf("error migration seeder: %v", err)
			return err
		}
	}

	if err := seed.PostSeed(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
		return err
	}

	return nil
}
