package utils

import (
	"github.com/Matheus1296/ProjectGoLang/domain"
	"os"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil{
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("dsn")

	db,err := gorm.Open("postgres",dsn)

	if err != nil{
		log.Fatalf("error connecting to database %v",err)
		panic(err)
	}

	//defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db
}