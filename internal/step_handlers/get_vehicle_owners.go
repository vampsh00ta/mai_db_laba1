package step_handlers

import (
	"TgDbMai/internal/response"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (sh StepHandler) CheckVehicleOwners(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите номер автомобиля",
	})
	bot.RegisterStepHandler(ctx, update, sh.checkVehicleOwnersResult, "")
}

func (sh StepHandler) checkVehicleOwnersResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	defer bot.UnregisterStepHandler(ctx, update)

	pts := update.Message.Text
	persons, err := sh.s.GetVehicleOwners(pts)
	if err != nil {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Что-то пошло не так"),
		})
	}
	//_, err = bot.SendDocument(ctx, &tgbotapi.SendDocumentParams{
	//	ChatID:      update.Message.Chat.ID,
	//	Document:       models.InputFileUpload{
	//		Data:dtps_file,
	//	},
	//})
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Владельцы автомобиля",
		ReplyMarkup: response.VehicleOwners(persons),
	})
}
