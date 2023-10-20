package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Main struct {
	*BotHandler
}

func NewMain(bot *tgbotapi.Bot, handler *BotHandler) {
	main := Main{handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GaishnikCommand,
		tgbotapi.MatchTypeExact,
		main.Gaishnik())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.GaiCommand,
		tgbotapi.MatchTypeExact,
		main.Gai())
}

//func NewMain(bot *tgbotapi.Bot){
//	main:=new(Main)
//	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
//		keyboard.CheckVehicleCommand,
//		tgbotapi.MatchTypeExact,
//		main.Gaishnik())
//	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
//		keyboard.AddParticipantDtpCommand,
//		tgbotapi.MatchTypeExact,
//		main.Gai())
//}

func (g Main) Gaishnik() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.back.keyboard = keyboard.Main()
		g.back.name = keyboard.MainСommand
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Меню действий сотрудника ГИБДД",
			ReplyMarkup: keyboard.Gaishnik(),
		})
	}
}
func (g Main) Gai() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.back.keyboard = keyboard.Main()
		g.back.name = keyboard.MainСommand
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Меню действий диспетчера ГИБДД",
			ReplyMarkup: keyboard.Gai(),
		})
	}
}
