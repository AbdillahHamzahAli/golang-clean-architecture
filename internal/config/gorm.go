package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabase() *gorm.DB {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error loading .env file")
			panic(err)
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db

}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = dbSql.Close()
}
