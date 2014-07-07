package main

import (
	"log"

	"github.com/capoferro/authosaurus/resources"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db gorm.DB

func main() {
	var err error

	dbPath := "./authosaurus.db"
	log.Printf("Opening database: " + dbPath)
	db, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Error connecting to the database: " + err.Error())
		log.Printf("No migrations performed.")
		return
	}

	log.Printf("Migrating...")
	db.AutoMigrate(resources.User{})
	log.Printf("Completed automigrate.")
}
