package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"TgDbMai/internal/step_handlers"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type MasterData struct {
	*BotHandler
	step *step_handlers.StepHandler
}

func NewMasterData(bot *tgbotapi.Bot, handler *BotHandler) {
	masterdata := MasterData{BotHandler: handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.AddCommand,
		tgbotapi.MatchTypeExact,
		masterdata.Add())

}
func (g *MasterData) Add() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.back.Set(update.Message.Chat.ID,
			&Back{name: keyboard.MasterDataCommand,
				keyboard: keyboard.MasterData(),
			},
		)

		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        keyboard.MasterDataCommand,
			ReplyMarkup: keyboard.Add(),
		})
	}
}
