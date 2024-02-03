package db

import (
	"book_inventory/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load file")
	}

	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)

	if err != nil {
		log.Fatal("error connecting to db")
	}

	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Books{})

	data := models.Books{}

	if db.Find(&data).RecordNotFound() {
		seederBooks(db)
	}
}

func seederBooks(db *gorm.DB) {
	data := []models.Books{{
		Title:       "Eternal Giants: Unveiling the Secrets of Dinosaurs",
		Author:      "Dr. Victoria Sinclair, Paleontologist Extraordinaire",
		Description: "Embark on a thrilling journey through time as Dr. Sinclair unravels the mysteries of dinosaurs and their incredible discoveries.",
		Stock:       10,
	}, {
		Title:       "Ingenious Ingenuity: A Chronicle of Engineering Marvels",
		Author:      "Professor Benjamin Inventorius",
		Description: "Dive into the world of awe-inspiring engineering feats and groundbreaking innovations that have shaped our modern civilization.",
		Stock:       5,
	}}

	for _, v := range data {
		db.Create(&v)
	}
}
