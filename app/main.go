package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("APP_NAME"))

}
