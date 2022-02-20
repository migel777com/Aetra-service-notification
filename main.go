package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"notificationBot/database"
	"notificationBot/logger"
	"notificationBot/tgbot"
)

func main() {
	botToken := "5162458517:AAGatmkeUkyJZhtqjNeWAsBhJ7bdlsXyNuw"

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