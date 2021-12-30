package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	if len(args) == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "no args")
		c.bot.Send(msg)
		return
	}

	idx, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "wrong args "+args)
		c.bot.Send(msg)
		log.Println("wrong args", args)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)
		log.Printf("fail to get product with idx %d: %s\n", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}
