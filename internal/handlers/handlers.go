package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"TgDbMai/internal/step_handlers"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type BotHandler struct {
	step *step_handlers.StepHandler
	back *Back
}

func New(bot *tgbotapi.Bot, step *step_handlers.StepHandler) {
	back := &Back{name: "", keyboard: nil}
	botHandler := &BotHandler{step, back}
	NewMain(bot, botHandler)
	NewGaishnik(bot, botHandler)
	NewGai(bot, botHandler)
	NewCheckVehicle(bot, botHandler)
	NewSpravki(bot, botHandler)
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		"/start", tgbotapi.MatchTypeExact, Start())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.BackCommand, tgbotapi.MatchTypeExact, back.undo())

}
func Start() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Главное меню",
			ReplyMarkup: keyboard.Main(),
		})
	}
}
