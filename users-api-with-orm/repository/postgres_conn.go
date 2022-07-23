package repository

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {

	// dbURL := "postgres://aplos_qa:Aplos2022@aplos-qa.c6n8utnxko28.ap-south-1.rds.amazonaws.com:5432/patil_db"
	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	dsn := "host=aplos-qa.c6n8utnxko28.ap-south-1.rds.amazonaws.com user=aplos_qa password=Aplos2022 dbname=patil_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return db
}
