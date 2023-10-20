package query_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Gai struct {
	*BotHandler
}

func NewGai(bot *tgbotapi.Bot, handler *BotHandler) {
	gai := Gai{handler}
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.RegVehicleCommand,
		tgbotapi.MatchTypeExact,
		gai.RegVehicle())
	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.DtpHappenCommand,
		tgbotapi.MatchTypeExact,
		gai.DtpHappen())
}

func (g Gai) DtpHappen() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		g.step.Dtp(ctx, bot, update)
	}
}

func (g Gai) RegVehicle() tgbotapi.HandlerFunc {
	return func(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   keyboard.RegVehicleCommand,
		})
	}
}

//func (b BotHandler) Gai(pattern string) tgbotapi.HandlerFunc {
//	kb := keyboard.Gai(pattern)
//	sh := b.step
//	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
//		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
//			CallbackQueryID: update.CallbackQuery.ID,
//			ShowAlert:       false,
//		})
//
//		switch update.CallbackQuery.Data {
//		case kb.CallbackData(keyboard.RegDtpKey):
//			sh.Dtp(ctx, b, update)
//		case kb.CallbackData(keyboard.RegVehicleKey):
//			fmt.Printf(kb.CallbackData(keyboard.RegVehicleKey))
//		case kb.CallbackData(keyboard.BackToMainKey):
//			fmt.Printf(kb.CallbackData(keyboard.BackToMainKey))
//		}
//	}
//}
