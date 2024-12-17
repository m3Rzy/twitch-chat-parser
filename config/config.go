package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var TOKEN = ""
var CHANNEL = ""

func GetConfigs() (string, string, string) {
	err := godotenv.Load("../.env") 
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	if username == "" {
		log.Fatal("USERNAME not set in .env file")
	}

	log.Println("Username:", username)
	log.Println("Channel:", CHANNEL)
	log.Println("Token:", TOKEN)

	return "oauth:" + TOKEN, username, "#" + CHANNEL
}
