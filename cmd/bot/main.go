package main

import (
	"log"
	"os"
	"yt_music_bot_go/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found!")
	}
}

func main() {

	bot, err := tgbotapi.NewBotAPI(getEnv("TG_TOKEN", ""))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv("TG_TOKEN"); exists {
		return value
	}
	return defaultVal
}
