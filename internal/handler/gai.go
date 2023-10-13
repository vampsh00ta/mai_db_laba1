package handler

import (
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"TgDbMai/internal/psql"
	"context"
	"fmt"
)

func GaiDtpName(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		Text:   "вв-те имя",
	})

	b.RegisterStepHandler(ctx, update, GaiDtpSurname, psql.Person{})
}
func GaiDtpSurname(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {

	data := b.GetStepData(ctx, update)
	fmt.Println(data)
	person := data.(psql.Person)
	person.Name = update.Message.Text
	b.RegisterStepHandler(ctx, update, GaiDtpOtchestvo, person)

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "вв-те фамилию",
	})
}
func GaiDtpOtchestvo(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	person := data.(psql.Person)
	person.Surname = update.Message.Text
	b.RegisterStepHandler(ctx, update, Final, person)

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "вв-те отество",
	})
}
func Final(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.UnregisterStepHandler(ctx, update)
	person := data.(psql.Person)
	person.Patronymic = update.Message.Text
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   person.Name + " " + person.Surname + " " + person.Patronymic,
	})
}
