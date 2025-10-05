package cmd

import (
	"log"
	"os"

	"github.com/AbdillahHamzahAli/golang-clean-architecture/db/migrations"
	"github.com/AbdillahHamzahAli/golang-clean-architecture/db/migrations/seeder"
	"gorm.io/gorm"
)

func Commands(db *gorm.DB) {
	migrate := false
	seed := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seed" {
			seed = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
	}

	if seed {
		if err := seeder.RunSeeder(db); err != nil {
			log.Fatalf("error migration seeder: %v", err)
		}
		log.Println("seeder completed successfully")
	}
}
