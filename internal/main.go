package main

import (
	"log"

	"github.com/Vico1993/turbo-memory/internal/bot"
	"github.com/Vico1993/turbo-memory/internal/cron"
	"github.com/Vico1993/turbo-memory/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	// Load the database
	database.Init()

	// Load cron
	cron.Init()

	// Start bot
	bot.Init()
}
