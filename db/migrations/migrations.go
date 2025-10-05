package migrations

import (
	entity2 "github.com/AbdillahHamzahAli/golang-clean-architecture/internal/domain/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if err := db.AutoMigrate(
		&entity2.User{},
		&entity2.Post{},
	); err != nil {
		return err
	}
	return nil
}
