package step_handlers

import (
	rep "TgDbMai/internal/repository"
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
	fmt.Println(tgIds, personIds)
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

}
