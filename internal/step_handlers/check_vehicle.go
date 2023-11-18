package step_handlers

import (
	rep "TgDbMai/internal/repository"
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

	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Дтп,в которых учавствовал автомобиль",
		ReplyMarkup: response.Dpts(dtps),
	})
	sh.log.Debug("CheckVehicle", "status", "good")

}

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

	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Владельцы автомобиля",
		ReplyMarkup: response.VehicleOwners(persons),
	})
	sh.log.Debug("CheckVehicleOwners", "status", "good")

}

func (sh StepHandler) VehicleDtps(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите номер автомобиля",
	})
	bot.RegisterStepHandler(ctx, update, sh.vehicleDtpsPts, "")
}
func (sh StepHandler) vehicleDtpsPts(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	defer bot.UnregisterStepHandler(ctx, update)

	pts := update.Message.Text
	dtps, err := sh.s.VehicleDpts(pts)
	if err != nil {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Что-то пошло не так"),
		})
	}

	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Дтп автомобиля",
		ReplyMarkup: response.Dpts(dtps),
	})
	sh.log.Debug("VehicleDtps", "status", "good")

}

func (sh StepHandler) VehicleInfo(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите номер автомобиля",
	})
	bot.RegisterStepHandler(ctx, update, sh.vehicleInfoPts, "")

}
func (sh StepHandler) vehicleInfoPts(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	defer bot.UnregisterStepHandler(ctx, update)

	pts := update.Message.Text
	vehicles, err := sh.s.GetVehicleByPts(pts)
	if err != nil {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Что-то пошло не так"),
		})
	}

	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Дтп автомобиля",
		ReplyMarkup: response.Vehicles([]*rep.Vehicle{vehicles}),
	})
	sh.log.Debug("VehicleInfo", "status", "good")

}
