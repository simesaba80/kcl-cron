package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBURL string

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBURL = os.Getenv("DBURL")
	fmt.Println("DBURL: ", DBURL)
}
