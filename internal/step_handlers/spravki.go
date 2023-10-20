package step_handlers

import (
	"TgDbMai/internal/response"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type PtsPerson struct {
	fio string
	pts string
}

func (sh StepHandler) IsPersonOwner(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите номер автомобиля",
	})
	bot.RegisterStepHandler(ctx, update, sh.isPersonOwnerPts, PtsPerson{})
}

func (sh StepHandler) isPersonOwnerPts(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	pts := update.Message.Text
	data := bot.GetStepData(ctx, update)
	pts_fio := data.(PtsPerson)
	pts_fio.pts = pts

	bot.RegisterStepHandler(ctx, update, sh.isPersonOwnerFio, pts_fio)
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите ФИО человека",
	})
	bot.RegisterStepHandler(ctx, update, sh.isPersonOwnerFio, pts_fio)

}
func (sh StepHandler) isPersonOwnerFio(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	data := bot.UnregisterStepHandler(ctx, update)

	fio := update.Message.Text
	pts_fio := data.(PtsPerson)
	pts_fio.fio = fio

	fmt.Println(pts_fio)
	res, err := sh.s.IsPersonOwner(pts_fio.pts, pts_fio.fio)
	if err != nil {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Что-то пошло не так"),
		})
		return
	}
	var result string
	if res {
		result = "Автомобиль принадлежит человеку"
	} else {
		result = "Автомобиль не принадлежит человеку"

	}
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   result,
	})
}
func (sh StepHandler) GetPersonsVehicles(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите  ФИО человека",
	})
	bot.RegisterStepHandler(ctx, update, sh.GetPersonsVehiclesResult, nil)
}
func (sh StepHandler) GetPersonsVehiclesResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	fio := update.Message.Text
	defer bot.UnregisterStepHandler(ctx, update)

	vehicles, err := sh.s.GetPersonsVehiclesAndDtps(fio)
	if err != nil {
		SendError(ctx, bot, update)
		return
	}
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введите  ФИО человека",
		ReplyMarkup: response.GetPersonsVehicles(vehicles),
	})
}
