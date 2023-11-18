package step_handlers

import (
	"TgDbMai/internal/keyboard"
	rep "TgDbMai/internal/repository"
	"TgDbMai/internal/response"
	"context"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strconv"
)

type PtsPerson struct {
	fio string
	pts string
}

func (sh StepHandler) GetCurrentDtp(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	user := sh.Auth.GetUser(update.Message.Chat.ID)
	dtp, err := sh.s.GetCurrentDtp(user.PersonId)
	if err != nil {
		SendError(ctx, bot, update)
		return
	}
	result := ""
	if dtp.Id == 0 {
		result = "В данный момент вы свободны 🙂"
		SendMessage(ctx, bot, update, result)
		return
	}
	result = "Текущее дтп"
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        result,
		ReplyMarkup: response.CurrentDtp(dtp),
	})
	sh.log.Debug("GetCurrentDtp", "status", "good")

}

func (sh StepHandler) AddDtpDescription(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Введите дополнение по дтп",
		ReplyMarkup: keyboard.DescriptionUpdate(),
	})
	b.RegisterStepHandler(ctx, update, sh.addDtpDescriptionResult, nil)

}
func (sh StepHandler) addDtpDescriptionResult(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	defer b.UnregisterStepHandler(ctx, update)

	text := update.Message.Text
	user := sh.Auth.GetUser(update.Message.Chat.ID)

	dtp, err := sh.s.GetCurrentDtp(user.PersonId)

	if err != nil {
		SendError(ctx, b, update)
		return
	}

	_, err = sh.s.AddDescriptionToDtp(dtp.Id, text)
	if err != nil {
		SendError(ctx, b, update)
		return
	}
	result := ""
	if text == rep.ClosedDtp {
		err := sh.s.CloseDtp(dtp.Id)
		if err != nil {
			SendError(ctx, b, update)
			return
		}
		result = rep.ClosedDtp
	} else {
		result = "Комментарий  добавлен"
	}

	b.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        result,
		ReplyMarkup: keyboard.Gaishnik(),
	})
	sh.log.Debug("AddDtpDescription", "status", "good")

}

func (sh StepHandler) AddParticipant(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите паспорт участника дтп",
	})

	bot.RegisterStepHandler(ctx, update, sh.addParticipantPassport, &rep.ParticipantOfDtp{})

}
func (sh StepHandler) addParticipantPassport(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	data := bot.GetStepData(ctx, update)
	participant := data.(*rep.ParticipantOfDtp)
	passport, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendError(ctx, bot, update)
		return
	}

	participant.Person.Passport = passport

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите автомобиль участника дтп",
	})

	bot.RegisterStepHandler(ctx, update, sh.addParticipantVehicle, participant)

}
func (sh StepHandler) addParticipantVehicle(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	data := bot.GetStepData(ctx, update)
	participant := data.(*rep.ParticipantOfDtp)
	pts := update.Message.Text

	participant.Vehicle.Pts = pts

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите номер закона,нарушенный участником дтп",
	})

	bot.RegisterStepHandler(ctx, update, sh.addParticipantRole, participant)

}
func (sh StepHandler) addParticipantRole(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	data := bot.GetStepData(ctx, update)
	participant := data.(*rep.ParticipantOfDtp)
	violation := update.Message.Text
	participant.Violation.LawNumber = violation

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите роль участника",
	})

	bot.RegisterStepHandler(ctx, update, sh.addParticipantViolation, participant)

}
func (sh StepHandler) addParticipantViolation(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	defer bot.UnregisterStepHandler(ctx, update)
	data := bot.GetStepData(ctx, update)
	participant := data.(*rep.ParticipantOfDtp)

	role := update.Message.Text

	participant.Role = role
	user := sh.Auth.GetUser(update.Message.Chat.ID)
	dtp, err := sh.s.GetCurrentDtp(user.PersonId)
	if err != nil {
		SendError(ctx, bot, update)
		return
	}
	_, err = sh.s.AddParticipant(dtp.Id, participant.Vehicle.Pts, participant.Person.Passport, participant.Role, &participant.Violation.LawNumber)
	if err != nil {
		SendError(ctx, bot, update)
		return
	}
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Участник дтп успешно создан",
	})
	sh.log.Debug("AddParticipant", "status", "good")

}

func (sh StepHandler) IssueFine(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите паспорт человека",
	})

	bot.RegisterStepHandler(ctx, update, sh.issueFinePassport, &rep.Person{})

}
func (sh StepHandler) issueFinePassport(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	data := bot.GetStepData(ctx, update)
	person := data.(*rep.Person)
	passport, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendError(ctx, bot, update)
		return
	}
	person.Passport = passport

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите размер штрафа",
	})
	bot.RegisterStepHandler(ctx, update, sh.issueFineAmount, person)

}
func (sh StepHandler) issueFineAmount(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	data := bot.GetStepData(ctx, update)
	person := data.(*rep.Person)
	amount, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendError(ctx, bot, update)
		return
	}
	person.Fine = append(person.Fine, &rep.Fine{Amount: amount})

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите причину штрафа",
	})
	bot.RegisterStepHandler(ctx, update, sh.issueFineReason, person)

}
func (sh StepHandler) issueFineReason(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	defer bot.UnregisterStepHandler(ctx, update)
	data := bot.GetStepData(ctx, update)
	person := data.(*rep.Person)
	reason := update.Message.Text

	person.Fine[0].Reason = reason

	_, err := sh.s.IssueFine(person.Passport, person.Fine[0].Amount, person.Fine[0].Reason)

	if err != nil {
		SendError(ctx, bot, update)
		return
	}
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Штраф успешно выписан",
	})
	sh.log.Debug("IssueFine", "status", "good")

}

func (sh StepHandler) CheckFines(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите паспорт человека",
	})

	bot.RegisterStepHandler(ctx, update, sh.checkFinesPassport, nil)

}
func (sh StepHandler) checkFinesPassport(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	defer bot.UnregisterStepHandler(ctx, update)
	passport, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		SendError(ctx, bot, update)
		return
	}

	person, err := sh.s.GetFines(passport)
	fines := person.Fine
	if err != nil {
		SendError(ctx, bot, update)
		return
	}
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Штраф успешно выписан",
		ReplyMarkup: response.GetFines(fines),
	})
	sh.log.Debug("CheckFines", "status", "good")

}
