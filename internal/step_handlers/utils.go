package step_handlers

import (
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SendMessage(ctx context.Context, b *tgbotapi.Bot, update *models.Update, text string) error {
	_, err := b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text,
	})
	return err
}
