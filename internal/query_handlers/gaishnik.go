package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Gaishnik struct {
	*BotHandler
}

func NewGaishnik(bot *tgbotapi.Bot, handler *BotHandler) {
	gaishnik := Gaishnik{handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.CheckVehicleCommand,
		tgbotapi.MatchTypeExact,
		gaishnik.CheckVehicle())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.AddParticipantDtpCommand,
		tgbotapi.MatchTypeExact,
		gaishnik.AddDtpParticipant())

}

//
//func (g Gaishnik) CheckVehicleOwner() tgbotapi.HandlerFunc {
//	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
//		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
//			ChatID: update.Message.Chat.ID,
//			Text:   keyboard.VehicleOwnerCommand,
//		})
//
//	}
//
//}

func (g Gaishnik) AddDtpParticipant() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   keyboard.AddParticipantDtpCommand,
		})

	}
}
func (g Gaishnik) CheckVehicle() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.back.keyboard = keyboard.Gaishnik()
		g.back.name = keyboard.GaishnikCommand
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        keyboard.CheckVehicleCommand,
			ReplyMarkup: keyboard.CheckVehicle(),
		})

	}
}
