package handler

import (
	"TgDbMai/internal/psql"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// insert into dtp (date,area,street,coords,category,metro)  values(now(),'string','string','string, string','string','string') returning id;
// select c.id  from crew as c join  gai on gai.id = c.gai_id
// where c.duty = true
// and c.id not  in (select crew_id from crew_dtp)
// and  gai.metro = ? limit 1;
// insert into  dtp_description(text, time,dtp_id) values('были вызваны сотрудники дтп',now(),?);
// insert into crew_dtp (dtp_id,crew_id) values(?,?);
func Dtp(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	bot.AnswerCallbackQuery(ctx, &tgbotapi.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		Text:   "Введите район дтп",
	})

	bot.RegisterStepHandler(ctx, update, DtpArea, psql.Dtp{})
}

func DtpArea(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(psql.Dtp)
	dtp.Area = update.Message.Text

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите улицу дтп",
	})

	b.RegisterStepHandler(ctx, update, DtpStreet, dtp)
}

func DtpStreet(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(psql.Dtp)
	dtp.Street = update.Message.Text

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите координаты дтп",
	})

	b.RegisterStepHandler(ctx, update, DtpCoords, dtp)
}

func DtpCoords(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(psql.Dtp)
	dtp.Coords = update.Message.Text

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите категорию дтп",
	})

	b.RegisterStepHandler(ctx, update, DtpCategory, dtp)
}
func DtpCategory(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(psql.Dtp)
	dtp.Category = update.Message.Text

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите метро дтп",
	})
	b.RegisterStepHandler(ctx, update, DtpMetro, dtp)

}
func DtpMetro(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(psql.Dtp)
	dtp.Metro = update.Message.Text

	DtpFinal(ctx, b, update)
}
func DtpFinal(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.UnregisterStepHandler(ctx, update)
	dtp := data.(psql.Dtp)
	dtp = dtp
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Сотрудники были успешно вызваны,ID ДТП: %d", 1),
	})

}
