package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	"fmt"
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
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetCurrentDtpCommand,
		tgbotapi.MatchTypeExact,
		gaishnik.GetCurrentDtp())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.AddChangesToDtpCommand,
		tgbotapi.MatchTypeExact,
		gaishnik.AddChangesToDtpCommand())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.IssueFineCommand,
		tgbotapi.MatchTypeExact,
		gaishnik.IssueFine())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.CheckFinesCommand,
		tgbotapi.MatchTypeExact,
		gaishnik.CheckFines())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.IsPersonOwnerCommand,
		tgbotapi.MatchTypeExact,
		gaishnik.CheckVehicleOwner())
}

func (g Gaishnik) CheckVehicleOwner() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.step.IsPersonOwner(ctx, b, update)

	}

}
func (g Gaishnik) IssueFine() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.step.IssueFine(ctx, b, update)

	}
}
func (g Gaishnik) CheckFines() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		fmt.Println("asd")
		g.step.CheckFines(ctx, b, update)

	}
}
func (g Gaishnik) AddDtpParticipant() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.step.AddParticipant(ctx, b, update)

	}
}
func (g Gaishnik) AddChangesToDtpCommand() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.step.AddDtpDescription(ctx, b, update)

	}
}
func (g Gaishnik) GetCurrentDtp() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.step.GetCurrentDtp(ctx, b, update)

	}
}
func (g Gaishnik) CheckVehicle() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.back.Set(update.Message.Chat.ID,
			&Back{name: keyboard.GaishnikCommand,
				keyboard: keyboard.Gaishnik(),
			},
		)

		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        keyboard.CheckVehicleCommand,
			ReplyMarkup: keyboard.CheckVehicle(),
		})

	}
}
