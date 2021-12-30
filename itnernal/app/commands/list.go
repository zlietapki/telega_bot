package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputText := "Here all products:\n"

	prods := c.productService.List()
	for _, p := range prods {
		outputText += p.Title
		outputText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputText)

	offsetText, _ := json.Marshal(CommandOffset{
		Offset: 11,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(offsetText)),
		),
	)

	c.bot.Send(msg)
}
