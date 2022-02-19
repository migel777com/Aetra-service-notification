package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"notificationBot/database"
	"notificationBot/logger"
	"notificationBot/tgbot"
	"os"
)

func main() {
	_ = godotenv.Load("globals.env")
	botToken := os.Getenv("BOT_TOKEN")

	myLogger := logger.RegisterLogger("logs.txt")
	defer myLogger.File.Close()

	db, err := database.OpenDB(&myLogger)
	defer db.Db.Close()

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil{
		log.Println("Error when creating a bot instance occurred", err)
	}

	tgbot.InitializeBot(bot, &db, &myLogger)
}