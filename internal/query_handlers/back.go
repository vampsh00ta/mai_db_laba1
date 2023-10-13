package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
)

func Back(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	switch update.CallbackQuery.Data {
	case "back_" + strconv.Itoa(keyboard.BackToMainKey):
		b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.ID,
			ReplyMarkup: keyboard.Main(),
		})

	}
}
