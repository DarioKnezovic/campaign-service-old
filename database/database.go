package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDatabase() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	_, err := gorm.Open("postgres", "postgres://"+username+":"+password+"@"+dbHost+"/"+dbName)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("WOHOOOOO")
}
