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
type BackSession struct {
	user map[int64]*Back
}

func (b BackSession) Get(userId int64) *Back {
	_, ok := b.user[userId]
	if !ok {
		return new(Back)
	}
	return b.user[userId]
}
func (b BackSession) Set(userId int64, back *Back) {
	b.user[userId] = back
}

func (b *BackSession) undo() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		back := b.Get(update.Message.Chat.ID)
		fmt.Println(b.user)
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        back.name,
			ReplyMarkup: back.keyboard,
		})
	}
}
func (back *Back) Exit() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	}
}
