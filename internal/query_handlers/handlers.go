package query_handlers

import (
	//"TgDbMai/internal/handler"
	"TgDbMai/internal/keyboard"
	"TgDbMai/internal/step_handlers"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
)

type BotHandler struct {
	step *step_handlers.StepHandler
	*tgbotapi.Bot
}

func (b BotHandler) Start(ctx context.Context, _ *tgbotapi.Bot, update *models.Update) {
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выберите",
		ReplyMarkup: keyboard.Main("").Markup(),
	})
}
func (b BotHandler) Back(pattern string) tgbotapi.HandlerFunc {
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		switch update.CallbackQuery.Data {
		case "back_" + strconv.Itoa(keyboard.BackToMainKey):
			b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
				ChatID:      update.CallbackQuery.Message.Chat.ID,
				MessageID:   update.CallbackQuery.Message.ID,
				ReplyMarkup: keyboard.Main("").Markup(),
			})
		case "back_" + strconv.Itoa(keyboard.BackToGaishnikKey):
			b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
				ChatID:      update.CallbackQuery.Message.Chat.ID,
				MessageID:   update.CallbackQuery.Message.ID,
				ReplyMarkup: keyboard.Gaishnik("").Markup(),
			})

		}
	}
}
func (b BotHandler) Gai(pattern string) tgbotapi.HandlerFunc {
	kb := keyboard.Gai(pattern)
	sh := b.step
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})

		switch update.CallbackQuery.Data {
		case kb.CallbackData(keyboard.RegDtpKey):
			sh.Dtp(ctx, b, update)
		case kb.CallbackData(keyboard.RegVehicleKey):
			fmt.Printf(kb.CallbackData(keyboard.RegVehicleKey))
		case kb.CallbackData(keyboard.BackToMainKey):
			fmt.Printf(kb.CallbackData(keyboard.BackToMainKey))
		}
	}
}
func (b BotHandler) Gaishnik(pattern string) tgbotapi.HandlerFunc {
	kb := keyboard.Gaishnik(pattern)
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		switch update.CallbackQuery.Data {
		case kb.CallbackData(keyboard.CheckVehicleOwnerKey):
			fmt.Printf(strconv.Itoa(keyboard.CheckVehicleOwnerKey))

		case kb.CallbackData(keyboard.AddDtpParticipantKey):
			fmt.Printf(strconv.Itoa(keyboard.AddDtpParticipantKey))

		case kb.CallbackData(keyboard.CheckVehicleKey):
			b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
				ChatID:      update.CallbackQuery.Message.Chat.ID,
				MessageID:   update.CallbackQuery.Message.ID,
				ReplyMarkup: keyboard.CheckVehicle("").Markup(),
			})

		}
	}
}
func (b BotHandler) Main(pattern string) tgbotapi.HandlerFunc {
	kb := keyboard.Main(pattern)
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		switch update.CallbackQuery.Data {
		case kb.CallbackData(keyboard.GaiKey):
			b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
				ChatID:      update.CallbackQuery.Message.Chat.ID,
				MessageID:   update.CallbackQuery.Message.ID,
				ReplyMarkup: keyboard.Gai("").Markup(),
			})
		case kb.CallbackData(keyboard.GaishnikKey):
			b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
				ChatID:      update.CallbackQuery.Message.Chat.ID,
				MessageID:   update.CallbackQuery.Message.ID,
				ReplyMarkup: keyboard.Gaishnik("").Markup(),
			})
		}
	}
}
func (b BotHandler) CheckVehicle(pattern string) tgbotapi.HandlerFunc {
	kb := keyboard.CheckVehicle(pattern)
	sh := b.step
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		switch update.CallbackQuery.Data {
		case kb.CallbackData(keyboard.CheckVehicleDtpsKey):
			sh.CheckVehicle(ctx, b, update)

		case kb.CallbackData(keyboard.CheckVehicleOwnerKey):
			sh.CheckVehicleOwners(ctx, b, update)

		}
	}
}
