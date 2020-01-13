package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init(){
	//Loading db creds from .env
	e := godotenv.Load()

	if e!=nil {
		fmt.Print(e)
	}
	//Getting Db creds
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	//Build connection string
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)

	if err!= nil {
		fmt.Println(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})

}
func GetDb() *gorm.DB{
	return db
}