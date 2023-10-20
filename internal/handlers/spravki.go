package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Spravki struct {
	handler *BotHandler
}

// IsPersonOwnerCommand      = "Проверить,принадлежит авто человеку"
// GetPersonsVehiclesCommand = "Вывести автомобили, принадлежащие человеку"
// GetPersonInfoCommand          = "Вывести данные человека"
// GetOfficersInfoCommand        = "ВЫвести данные сотрудника ГИБДД"
// GetDtpsInfoNearMetroCommand   = "Вывести ДТП, произошедшие у конкретного метро"
// GetDtpsInfoRadiusMetroCommand
func NewSpravki(bot *tgbotapi.Bot, handler *BotHandler) {
	spravki := Spravki{handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.IsPersonOwnerCommand,
		tgbotapi.MatchTypeExact,
		spravki.IsPersonOwner())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetPersonsVehiclesCommand,
		tgbotapi.MatchTypeExact,
		spravki.GetPersonsVehicles())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetPersonInfoCommand,
		tgbotapi.MatchTypeExact,
		spravki.GetPersonInfo())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetOfficersInfoCommand,
		tgbotapi.MatchTypeExact,
		spravki.GetOfficersInfoCommand())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetDtpsInfoNearMetroCommand,
		tgbotapi.MatchTypeExact,
		spravki.GetDtpsInfoNearMetro())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GetDtpsInfoRadiusMetroCommand,
		tgbotapi.MatchTypeExact,
		spravki.GetDtpsInfoRadiusMetro())

}
func (s Spravki) IsPersonOwner() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.handler.step.IsPersonOwner(ctx, bot, update)

	}
}

// добабив вывод дтп, в которых был автомобиль
func (s Spravki) GetPersonsVehicles() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		s.handler.step.GetPersonsVehicles(ctx, bot, update)
	}
}
func (s Spravki) GetPersonInfo() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   keyboard.GetPersonInfoCommand,
		})
	}
}
func (s Spravki) GetOfficersInfoCommand() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   keyboard.GetOfficersInfoCommand,
		})
	}
}
func (s Spravki) GetDtpsInfoNearMetro() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   keyboard.GetDtpsInfoNearMetroCommand,
		})
	}
}
func (s Spravki) GetDtpsInfoRadiusMetro() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   keyboard.GetDtpsInfoRadiusMetroCommand,
		})
	}
}
