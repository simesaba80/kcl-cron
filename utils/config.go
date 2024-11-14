package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DBURL1 string
var DBURL2 string

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBURL1 = os.Getenv("DBURL1")
	DBURL2 = os.Getenv("DBURL2")
	fmt.Println("DBURL: ", DBURL1)
	fmt.Println("DBURL: ", DBURL2)
}
