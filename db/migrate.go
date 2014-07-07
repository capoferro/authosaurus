package main

import (
	"os"
	"log"

	"github.com/capoferro/authosaurus/resources"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db gorm.DB

func main() {
	var err error
	var dbFilename string
	if os.Getenv("TEST") == "true" {
		dbFilename = "./authosaurus_test.db"
	} else {
		dbFilename = "./authosaurus.db"
	}
	log.Printf("Opening database: " + dbFilename)
	db, err = gorm.Open("sqlite3", dbFilename)
	if err != nil {
		log.Printf("Error connecting to the database: " + err.Error())
		log.Printf("No migrations performed.")
		return
	}

	log.Printf("Migrating...")
	db.AutoMigrate(resources.User{})
	log.Printf("Completed automigrate.")
}
