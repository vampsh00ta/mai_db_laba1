package step_handlers

import (
	"TgDbMai/internal/keyboard"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (sh StepHandler) Login(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {

	_, err := bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Введите ваще ФИО",
	})
	if err != nil {
		sh.log.Error(err.Error())
		SendError(ctx, bot, update)
		return
	}
	bot.RegisterStepHandler(ctx, update, sh.loginResult, nil)
}
func (sh StepHandler) loginResult(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	defer bot.UnregisterStepHandler(ctx, update)
	fio := update.Message.Text
	officer, _, err := sh.s.GetOfficersInfo(fio)
	if err != nil {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Что-то пошло не так"),
		})
		return
	}
	result := ""
	if officer != nil {
		me, err := bot.GetMe(ctx)
		if err != nil {
			bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("Что-то пошло не так"),
			})
		}
		userTgId := me.ID

		sh.Auth.LogIn(userTgId, officer.PersonId, 2)
		result = "Вы успешно авторизириовались"

	} else {
		result = "Неправильные данные"
	}
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        result,
		ReplyMarkup: keyboard.Gaishnik(),
	})

}

func (sh StepHandler) Logout(ctx context.Context, bot *tgbotapi.Bot, update *models.Update) {
	me, err := bot.GetMe(ctx)
	if err != nil {
		bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Что-то пошло не так"),
		})
		return
	}
	userTgId := me.ID
	sh.Auth.LogOut(userTgId)
	bot.SendMessage(ctx, &tgbotapi.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Вы успешно вышли из аккаунта"),
	})

}
