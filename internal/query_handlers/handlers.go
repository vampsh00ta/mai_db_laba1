package query_handlers

import (
	"TgDbMai/internal/handler"
	"TgDbMai/internal/keyboard"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
)

func NewStart(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выберите",
		ReplyMarkup: keyboard.NewMain("").Markup(),
	})
}
func NewBack(pattern string) tgbotapi.HandlerFunc {
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
				ReplyMarkup: keyboard.NewMain("").Markup(),
			})
		case "back_" + strconv.Itoa(keyboard.BackToGaishnikKey):
			b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
				ChatID:      update.CallbackQuery.Message.Chat.ID,
				MessageID:   update.CallbackQuery.Message.ID,
				ReplyMarkup: keyboard.NewGaishnik("").Markup(),
			})

		}
	}
}
func NewGai(pattern string) tgbotapi.HandlerFunc {
	kb := keyboard.NewGai(pattern)
	return func(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
		b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})

		switch update.CallbackQuery.Data {
		case kb.CallbackData(keyboard.RegDtpKey):
			handler.Dtp(ctx, b, update)
		case kb.CallbackData(keyboard.RegVehicleKey):
			fmt.Printf(kb.CallbackData(keyboard.RegVehicleKey))
		case kb.CallbackData(keyboard.BackToMainKey):
			fmt.Printf(kb.CallbackData(keyboard.BackToMainKey))
		}
	}
}
func NewGaishnik(pattern string) tgbotapi.HandlerFunc {
	kb := keyboard.NewGaishnik(pattern)
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
				ReplyMarkup: keyboard.NewCheckVehicle("").Markup(),
			})

		}
	}
}
func NewMain(pattern string) tgbotapi.HandlerFunc {
	kb := keyboard.NewMain(pattern)
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
				ReplyMarkup: keyboard.NewGai("").Markup(),
			})
		case kb.CallbackData(keyboard.GaishnikKey):
			b.EditMessageReplyMarkup(ctx, &tgbotapi.EditMessageReplyMarkupParams{
				ChatID:      update.CallbackQuery.Message.Chat.ID,
				MessageID:   update.CallbackQuery.Message.ID,
				ReplyMarkup: keyboard.NewGaishnik("").Markup(),
			})
		}
	}
}
func NewVehicle(pattern string) tgbotapi.HandlerFunc {
	kb := keyboard.NewCheckVehicle(pattern)

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
			fmt.Printf(strconv.Itoa(keyboard.CheckVehicleDtpsKey))

		case kb.CallbackData(keyboard.CheckVehicleOwnerKey):
			fmt.Printf(strconv.Itoa(keyboard.CheckVehicleOwnerKey))

		}
	}
}
