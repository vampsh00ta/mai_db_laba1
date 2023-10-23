package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type MasterData struct {
	handler *BotHandler
	//step *step_handlers.StepHandler
}

func NewMasterData(bot *tgbotapi.Bot, handler *BotHandler) {
	masterdata := MasterData{handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.AddCommand,
		tgbotapi.MatchTypeExact,
		masterdata.Add())

}
func (g *MasterData) Add() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.handler.back.keyboard = keyboard.MasterData()
		g.handler.back.name = keyboard.MainСommand
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        keyboard.MasterDataCommand,
			ReplyMarkup: keyboard.Add(),
		})
	}
}
