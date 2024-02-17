package datastore

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

func Initalize() error {
	godotenv.Load()
	dbURI := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("POSTGIS_HOST"),
		os.Getenv("POSTGIS_USER"),
		os.Getenv("POSTGIS_DATABASE"),
		os.Getenv("POSTGIS_PASSWORD"))
	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println("couldn't connect to database")
		return err
	}
	DB = conn
	DB = DB.Debug()
	return nil
}
