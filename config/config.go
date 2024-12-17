package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var TOKEN = ""

func GetConfigs() (string, string, string) {
	err := godotenv.Load("../.env") 
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	if username == "" {
		log.Fatal("USERNAME not set in .env file")
	}

	channel := "#" + os.Getenv("CHANNEL")
	if channel == "#" {
		log.Fatal("CHANNEL not set in .env file")
	}

	log.Println("Username:", username)
	log.Println("Channel:", channel)
	log.Println("Token:", TOKEN)

	return "oauth:" + TOKEN, username, channel
}
