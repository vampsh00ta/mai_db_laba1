package step_handlers

import (
	rep "TgDbMai/internal/repository"
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

	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите номер автомобиля",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	bot.RegisterStepHandler(ctx, update, sh.isPersonOwnerPts, PtsPerson{})
}

func (sh StepHandler) isPersonOwnerPts(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	pts := update.Message.Text
	data := bot.GetStepData(ctx, update)
	pts_fio := data.(PtsPerson)
	pts_fio.pts = pts

	bot.RegisterStepHandler(ctx, update, sh.isPersonOwnerFio, pts_fio)
	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите ФИО человека",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	bot.RegisterStepHandler(ctx, update, sh.isPersonOwnerFio, pts_fio)

}
func (sh StepHandler) isPersonOwnerFio(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	data := bot.UnregisterStepHandler(ctx, update)

	fio := update.Message.Text
	pts_fio := data.(PtsPerson)
	pts_fio.fio = fio

	vehicle, err := sh.s.GetPersonVehicleByPts(pts_fio.pts, pts_fio.fio)
	if err != nil {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Что-то пошло не так"),
		})
		return
	}
	var result string
	if vehicle.Id != 0 {
		result = "Автомобиль принадлежит человеку"
	} else {
		result = "Автомобиль не принадлежит человеку"

	}
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   result,
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	sh.log.Debug("IsPersonOwner", "status", "good")

}
func (sh StepHandler) GetPersonsVehicles(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите  ФИО человека",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	bot.RegisterStepHandler(ctx, update, sh.getPersonsVehiclesResult, nil)
}
func (sh StepHandler) getPersonsVehiclesResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	fio := update.Message.Text
	defer bot.UnregisterStepHandler(ctx, update)

	vehicles, err := sh.s.GetPersonsVehiclesAndDtps(fio)
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}

	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введите  ФИО человека",
		ReplyMarkup: response.GetPersonsVehicles(vehicles),
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	sh.log.Debug("GetPersonsVehicles", "status", "good")

}

func (sh StepHandler) GetPersonInfoPassport(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите паспорт",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}

	bot.RegisterStepHandler(ctx, update, sh.getPersonInfoPassportResult, PtsPerson{})
}

func (sh StepHandler) getPersonInfoPassportResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	defer bot.UnregisterStepHandler(ctx, update)
	passport := update.Message.Text
	persons, err := sh.s.GetPersonInfoPassport(passport)
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Результат",
		ReplyMarkup: response.GetPersonInfo(persons),
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	sh.log.Debug("GetPersonInfoPassport", "status", "good")

}

func (sh StepHandler) GetPersonInfoFIO(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите ФИО",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	bot.RegisterStepHandler(ctx, update, sh.getPersonInfoFIOResult, PtsPerson{})
}

func (sh StepHandler) getPersonInfoFIOResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	defer bot.UnregisterStepHandler(ctx, update)
	fio := update.Message.Text
	persons, err := sh.s.GetPersonInfoFIO(fio)
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Результат",
		ReplyMarkup: response.GetPersonInfo(persons),
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	sh.log.Debug("GetPersonInfoFIO", "status", "good")

}

func (sh StepHandler) GetOfficersInfo(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите  ФИО сотрудника",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	bot.RegisterStepHandler(ctx, update, sh.getOfficersResult, nil)
}
func (sh StepHandler) getOfficersResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	fio := update.Message.Text
	defer bot.UnregisterStepHandler(ctx, update)

	officer, dtps, err := sh.s.GetOfficersInfo(fio)
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	fmt.Println(officer.Person)
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Информация от сотрудинке",
		ReplyMarkup: response.GetOfficerInfo([]*rep.PoliceOfficer{officer}),
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Дтп сотрудника",
		ReplyMarkup: response.Dpts(dtps),
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Взводы сотрудника",
		ReplyMarkup: response.Сrew([]*rep.Crew{&rep.Crew{}}),
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	sh.log.Debug("GetOfficersInfo", "status", "good")
}

func (sh StepHandler) GetDtpsInfoNearArea(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите  район ",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	bot.RegisterStepHandler(ctx, update, sh.getDtpsInfoNearAreaResult, nil)
}

func (sh StepHandler) getDtpsInfoNearAreaResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	area := update.Message.Text
	defer bot.UnregisterStepHandler(ctx, update)
	dtps, err := sh.s.GetDtpsInfoNearArea(area)
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Дтп",
		ReplyMarkup: response.Dpts(dtps),
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}

	sh.log.Debug("GetOfficersInfo", "status", "good")
}
