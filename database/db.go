package database

import (
	"final_project/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func SetupDB() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error reading .env file : ", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	} else {
		fmt.Println("successfully connected to database")
	}

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB(db *gorm.DB) {
	var dbSql, err = db.DB()
	if err != nil {
		log.Fatal("failed to close database connection : ", err)
	}

	dbSql.Close()
}
