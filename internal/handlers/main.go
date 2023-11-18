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

	bot.RegisterHandler(tgbotapi.HandlerTypeMessageText,
		keyboard.ExitCommand,
		tgbotapi.MatchTypeExact,
		main.Exit())
}

func (g *Main) Gaishnik() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.back.Set(update.Message.Chat.ID,
			&Back{
				name:     keyboard.MainСommand,
				keyboard: keyboard.Main(),
			},
		)

		logged := g.step.Auth.IsLogged(update.Message.Chat.ID)

		if !logged {
			g.step.Login(ctx, b, update)
		} else {
			b.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID:      update.Message.Chat.ID,
				Text:        keyboard.GaishnikCommand,
				ReplyMarkup: keyboard.Gaishnik(),
			})
		}
		g.step.Producer.SendMsg("calls", keyboard.GaishnikCommand)

	}
}
func (g *Main) Gai() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.back.Set(update.Message.Chat.ID,
			&Back{
				name:     keyboard.MainСommand,
				keyboard: keyboard.Main(),
			},
		)
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        keyboard.GaiCommand,
			ReplyMarkup: keyboard.Gai(),
		})
		g.step.Producer.SendMsg("calls", keyboard.GaiCommand)

	}
}

func (g Main) Exit() tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		g.step.Logout(ctx, b, update)
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        keyboard.MainСommand,
			ReplyMarkup: keyboard.Main(),
		})
		g.step.Producer.SendMsg("calls", keyboard.ExitCommand)

	}
}
