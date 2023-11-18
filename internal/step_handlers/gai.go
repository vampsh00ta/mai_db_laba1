package step_handlers

import (
	rep "TgDbMai/internal/repository"
	"TgDbMai/internal/response"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
)

type dtpOfficerCount struct {
	rep.Dtp
	count int
}
type CoordsRadius struct {
	Coords string
	Radius int
}

func (sh StepHandler) Dtp(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {

	err := SendMessage(ctx, b, update, "Введите район дтп")
	if err != nil {
		SendError(ctx, b, update)
		return
	}

	b.RegisterStepHandler(ctx, update, sh.dtpArea, dtpOfficerCount{})
}

func (sh StepHandler) dtpArea(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(dtpOfficerCount)
	dtp.Area = update.Message.Text
	err := SendMessage(ctx, b, update, "Введите улицу дтп")

	if err != nil {
		SendError(ctx, b, update)
		return
	}

	b.RegisterStepHandler(ctx, update, sh.dtpStreet, dtp)
}

func (sh StepHandler) dtpStreet(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(dtpOfficerCount)
	dtp.Street = update.Message.Text
	err := SendMessage(ctx, b, update, "Введите координаты дтп")

	if err != nil {
		SendError(ctx, b, update)
		return
	}

	b.RegisterStepHandler(ctx, update, sh.dtpCoords, dtp)
}

func (sh StepHandler) dtpCoords(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(dtpOfficerCount)
	dtp.Coords = update.Message.Text
	err := SendMessage(ctx, b, update, "Введите категорию дтп")

	if err != nil {
		SendError(ctx, b, update)
		return
	}

	b.RegisterStepHandler(ctx, update, sh.dtpCategory, dtp)
}
func (sh StepHandler) dtpCategory(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(dtpOfficerCount)
	dtp.Category = update.Message.Text

	err := SendMessage(ctx, b, update, "Введите метро дтп")

	if err != nil {
		SendError(ctx, b, update)
		return
	}
	b.RegisterStepHandler(ctx, update, sh.dtpMetro, dtp)

}
func (sh StepHandler) dtpMetro(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(dtpOfficerCount)
	dtp.Metro = update.Message.Text

	err := SendMessage(ctx, b, update, "Введите число сотрудников")

	if err != nil {
		SendError(ctx, b, update)
		return
	}
	b.RegisterStepHandler(ctx, update, sh.dtpCount, dtp)

	//sh.dtpResult(ctx, b, update)
}

func (sh StepHandler) dtpCount(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	data := b.GetStepData(ctx, update)
	dtp := data.(dtpOfficerCount)
	count, err := strconv.Atoi(update.Message.Text)
	dtp.count = count

	defer b.UnregisterStepHandler(ctx, update)

	newDtp, err := sh.s.RegDtp(&dtp.Dtp, dtp.count)
	if err != nil {
		SendError(ctx, b, update)
		return
	}

	err = SendMessage(ctx, b, update,
		fmt.Sprintf("Сотрудники были успешно вызваны,ID ДТП: %d", newDtp.Id))
	if err != nil {

		SendError(ctx, b, update)
		return
	}

	var personIds []int
	for _, crew := range newDtp.Crews {
		for _, officer := range crew.PoliceOfficers {
			personIds = append(personIds, officer.PersonId)
		}
	}
	tgIds := sh.Auth.GetTgIdsByPersonId(personIds...)
	for _, id := range tgIds {
		if id == 0 {
			continue
		}
		_, err := b.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: id,
			Text:   fmt.Sprintf("Произошло дтп по координатам %s", dtp.Coords),
		})

		if err != nil {
			SendError(ctx, b, update)
			return
		}
	}
	sh.log.Debug("Dtp", "status", "good")

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
		SendError(ctx, bot, update)
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

	vehicles, err := sh.s.GetPersonsVehicles(fio)
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}

	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введите  ФИО человека",
		ReplyMarkup: response.Vehicles(vehicles),
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
	_, err = bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Информация от сотрудинке",
		ReplyMarkup: response.GetOfficerInfo(officer),
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
		ReplyMarkup: response.Сrew(officer.PoliceOfficer[0].Crews),
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

	sh.log.Debug("GetDtpsInfoNearArea", "status", "good")
}

func (sh StepHandler) GetDtpsInfoRadius(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите  координаты района ",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	bot.RegisterStepHandler(ctx, update, sh.getDtpsInfoRadiusCoords, CoordsRadius{})
}

func (sh StepHandler) getDtpsInfoRadiusCoords(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	data := bot.GetStepData(ctx, update)
	c := data.(CoordsRadius)
	coords := update.Message.Text
	c.Coords = coords

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите радиус",
	})

	bot.RegisterStepHandler(ctx, update, sh.getDtpsInfoRadius, c)
}
func (sh StepHandler) getDtpsInfoRadius(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	radius, err := strconv.Atoi(update.Message.Text)
	data := bot.GetStepData(ctx, update)
	c := data.(CoordsRadius)
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	defer bot.UnregisterStepHandler(ctx, update)
	dtps, err := sh.s.GetDtpsInfoRadius(radius, c.Coords)
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

	sh.log.Debug("GetDtpsInfoRadius", "status", "good")

}
