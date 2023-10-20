package step_handlers

import (
	"TgDbMai/internal/response"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (sh StepHandler) CheckVehicle(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите номер автомобиля",
	})
	bot.RegisterStepHandler(ctx, update, sh.checkVehicleResult, "")
}

func (sh StepHandler) checkVehicleResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	defer bot.UnregisterStepHandler(ctx, update)

	pts := update.Message.Text
	dtps, err := sh.s.VehicleDpts(pts)

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
		Text:        "Дтп,в которых учавствовал автомобиль",
		ReplyMarkup: response.VehicleDpts(dtps),
	})
}
