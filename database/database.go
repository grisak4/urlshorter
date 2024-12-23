package database

import (
	"fmt"
	"log"
	"urlshorter/config"
	"urlshorter/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	var err error
	confDb := config.GetDatabaseConf()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		confDb.Host, confDb.Port, confDb.User, confDb.Password, confDb.DBName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}

	if err = db.AutoMigrate(&models.ShortUrl{}); err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}

	log.Println("[DB] Successfully started!")
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if sqlDB, err := db.DB(); err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	} else {
		sqlDB.Close()
	}
}
