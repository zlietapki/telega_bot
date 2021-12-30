package commands

import (
	"encoding/json"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zlietapki/telega_bot/itnernal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

type CommandOffset struct {
	Offset int `json:"offset"`
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HandlerUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Println("panic recoveres", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		var parsedData CommandOffset
		_ = json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)

		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Offset "+strconv.Itoa(parsedData.Offset),
		)
		c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
