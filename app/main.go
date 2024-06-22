package main

import (
	"fmt"
	"golang-crud/app/src/controller/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const APP_NAME = "APP_NAME"

// loadEnv loads environment variables from a .env file.
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// getEnv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
func getEnv(key string) string {
	return os.Getenv(key)
}

func main() {
	loadEnv()
	appName := getEnv(APP_NAME)
	if appName == "" {
		log.Fatal("APP_NAME environment variable not set")
	}
	fmt.Println(appName)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}