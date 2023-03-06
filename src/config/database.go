package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	connectionBuilder := fmt.Sprintf("host=%s password=%s dbname=%s user=%s port=%s sslmode=disable", host, password, name, user, port)
	fmt.Println(connectionBuilder)
	DB, err = gorm.Open(postgres.Open(connectionBuilder), &gorm.Config{})
	if err != nil {
		fmt.Println("error in connecting database")
	} else {
		fmt.Println("database connect successfully ....")
	}
}
