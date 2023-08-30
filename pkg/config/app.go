package config

import (
	"os"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	url := os.Getenv("URL")
	port := os.Getenv("PORT")
	// "root:qwerty1234?charset=utf8&parseTime=True&loc=Local"
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := user + ":" + password + "@tcp(" + url + port + ")/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := `root:qwerty1234@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local`
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}