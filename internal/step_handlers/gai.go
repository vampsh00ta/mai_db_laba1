package step_handlers

import (
	rep "TgDbMai/internal/repository"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (sh StepHandler) Dtp(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите район дтп",
	})
	bot.RegisterStepHandler(ctx, update, sh.dtpArea, rep.Dtp{})
}

func (sh StepHandler) dtpArea(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(rep.Dtp)
	dtp.Area = update.Message.Text

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите улицу дтп",
	})

	b.RegisterStepHandler(ctx, update, sh.dtpStreet, dtp)
}

func (sh StepHandler) dtpStreet(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(rep.Dtp)
	dtp.Street = update.Message.Text

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите координаты дтп",
	})

	b.RegisterStepHandler(ctx, update, sh.dtpCoords, dtp)
}

func (sh StepHandler) dtpCoords(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(rep.Dtp)
	dtp.Coords = update.Message.Text

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите категорию дтп",
	})

	b.RegisterStepHandler(ctx, update, sh.dtpCategory, dtp)
}
func (sh StepHandler) dtpCategory(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(rep.Dtp)
	dtp.Category = update.Message.Text

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите метро дтп",
	})
	b.RegisterStepHandler(ctx, update, sh.dtpMetro, dtp)

}
func (sh StepHandler) dtpMetro(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(rep.Dtp)
	dtp.Metro = update.Message.Text

	sh.dtpResult(ctx, b, update)
}
func (sh StepHandler) dtpResult(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.UnregisterStepHandler(ctx, update)
	dtp := data.(rep.Dtp)
	id, err := sh.s.RegDtp(&dtp)
	if err != nil {
		b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Что-то пошло не так"),
		})
	}
	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Сотрудники были успешно вызваны,ID ДТП: %d", id),
	})

}
