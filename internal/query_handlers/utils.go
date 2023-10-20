package query_handlers

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Back struct {
	keyboard *models.ReplyKeyboardMarkup
	name     string
}

func (back *Back) undo() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		fmt.Println(back.name, back.keyboard)
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        back.name,
			ReplyMarkup: back.keyboard,
		})
	}
}
