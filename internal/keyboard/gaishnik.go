package keyboard

import (
	"github.com/go-telegram/bot/models"
)

func Gaishnik() *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Проверить владельца автомобиля", CallbackData: "gaishnik_auto_owner"},
			}, {
				{Text: "Добавить участника дтп", CallbackData: "gaishnik_add_participant"},
			},
			{
				{Text: "Проверить автомобиль", CallbackData: "gaishnik_check_auto"},
			},
			//{
			//	{Text: "Выписать штраф", CallbackData: "gaishnik_test"},
			//},
			{
				{Text: "Назад", CallbackData: "back_main"},
			},
		},
	}
	return kb
}
