package main

import (
	"log"
	"os"

	"github.com/zlietapki/telega_bot/itnernal/app/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/zlietapki/telega_bot/itnernal/service/product"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.UpdateConfig{
		Timeout: 60,
		Offset:  0,
	}

	updates := bot.GetUpdatesChan(updateConfig)

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandlerUpdate(update)
	}
}
