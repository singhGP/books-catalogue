package db

import (
	"log"

	"github.com/KaurKaranjot/bookslib/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "host=localhost port=5432 dbname=library user=postgres password=root"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Book{})
	return db
}
